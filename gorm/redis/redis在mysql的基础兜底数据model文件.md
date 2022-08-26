```golang
package redis

import "time"

// DDL
//CREATE TABLE `jx_redis_storage` (
//`id` int(11) NOT NULL AUTO_INCREMENT,
//`parent_key_id` int(11) NOT NULL DEFAULT '0' COMMENT '0为根节点数据，父级key对应的ID，如hset的某个field的parentKey为hset key在本表中的ID',
//`key` varchar(150) NOT NULL COMMENT 'redis存储的最小单位唯一标识值，如hset中的field值',
//`value` varchar(255) DEFAULT '' COMMENT 'redis存储的value值',
//`created_at` datetime DEFAULT NULL,
//`updated_at` datetime DEFAULT NULL,
//`deleted_at` datetime DEFAULT NULL,
//PRIMARY KEY (`id`),
//UNIQUE KEY `pid_key_unique_id` (`parent_key_id`,`key`) USING BTREE COMMENT '父节点id与key构成唯一'
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='本表为专门存储redis配置的兜底数据，仅支持基础数据类型，遍历深度最大为2。\n';

type Node struct {
 Key string `json:"key"`
 Value string `json:"value"`
 ParentNode *Node `json:"parent_node"`
 ChildNodes []*Node `json:"child_nodes"`
}

const (
 RootNodeKeyId = 0 //根节点数据默认值
)

type JxRedisStorage struct {
 ID int `gorm:"primaryKey;column:id;type:int(11);not null"`
 ParentKeyId int `gorm:"column:parent_key_id;type:int(11)"`
 Key string `gorm:"column:key;type:varchar(150)"`
 Value string `gorm:"column:value;type:varchar(255)"`
 CreatedAt *time.Time `gorm:"column:created_at;type:datetime"`
 UpdatedAt *time.Time `gorm:"column:updated_at;type:datetime"`
 DeletedAt *time.Time `gorm:"column:deleted_at;type:datetime"`
}

// TableName get sql table name.获取数据库表名
func (m *JxRedisStorage) TableName() string {
 return "jx_redis_storage"
}

func (m *JxRedisStorage) CommonQry(whereSql interface{}, args ...interface{}) ([]*JxRedisStorage, error) {
 temp := make([]*JxRedisStorage, 0)
 e := global.JX_DB.Model(new(JxRedisStorage)).Where(whereSql, args...).Find(&temp).Error
 return temp, e
}

// GetByKey 获取一个对应redis中的key：只遍历到深度2
func (m *JxRedisStorage) GetByKey(key string) (*Node, error) {
 root := &JxRedisStorage{}
 e := global.JX_DB.Model(new(JxRedisStorage)).Where("key = ? and parent_key_id = ?", key, RootNodeKeyId).First(&root).Error
 if e != nil {
 return nil, e
 }
 rootNode := &Node{
 Key: root.Key,
 Value: root.Value,
 ParentNode: nil,
 ChildNodes: []*Node{},
 }
 childNodes, err := m.CommonQry("parent_key_id = ?", root.ID)
 if err != nil {
 return nil, err
 }
 for i := range childNodes {
 rootNode.ChildNodes = append(rootNode.ChildNodes, &Node{
 Key: childNodes[i].Key,
 Value: childNodes[i].Value,
 ParentNode: rootNode,
 ChildNodes: []*Node{},
 })
 }
 return rootNode, nil
}

//从新增family节点动作事务抽象出来的方法体
func (m *JxRedisStorage) setKeyByFamilyNodeTransactionFunc(rootNode *Node) func(tx *gorm.DB) error {
 return func(tx *gorm.DB) error {
 toInsertRoot := &JxRedisStorage{
 ParentKeyId: RootNodeKeyId,
 Key: rootNode.Key,
 Value: rootNode.Value,
 }
 if e := tx.Create(toInsertRoot).Error; e != nil {
 return e
 }
 toInsertChildNodes := []*JxRedisStorage{}
 for i := range rootNode.ChildNodes {
 toInsertChildNodes = append(toInsertChildNodes, &JxRedisStorage{
 ParentKeyId: toInsertRoot.ID,
 Key: rootNode.ChildNodes[i].Key,
 Value: rootNode.ChildNodes[i].Value,
 })
 }
 return tx.Create(toInsertChildNodes).Error
 }
}

// SetKeyByFamilyNode 保存根节点及其子节点们的数据，即redis的一个完整key
func (m *JxRedisStorage) SetKeyByFamilyNode(rootNode *Node) (bool, error) {
 //解析根节点，遍历出所有节点及其子节点数据，一致地新增, 同样只遍历到深度2
 e := global.JX_DB.Transaction(m.setKeyByFamilyNodeTransactionFunc(rootNode))
 if e != nil {
 return false, e
 }
 return true, nil
}

// UpdateKeyByFamily 将原来的所有节点数据全部delete掉，然后新增
func (m *JxRedisStorage) UpdateKeyByFamily(familyNode *Node) (bool, error) {
 e := global.JX_DB.Transaction(func(tx *gorm.DB) error {
 rootNodeModel := &JxRedisStorage{}
 if err := tx.Model(new(JxRedisStorage)).Where("parent_key_id = ? and key = ?", RootNodeKeyId, familyNode.Key).First(&rootNodeModel).Error; err != nil {
 return err
 }
 if err := tx.Model(new(JxRedisStorage)).Where("(parent_key_id = ? and key = ?) or (parent_key_id = ?)", RootNodeKeyId, familyNode.Key, rootNodeModel.ID).Delete(m).Error; err != nil {
 return err
 }
 if err := m.setKeyByFamilyNodeTransactionFunc(familyNode)(tx); err != nil {
 return err
 }
 return nil
 })
 if e != nil {
 return false, e
 }
 return true, nil
}

// SaveKeyByFamilyNode 新增或修改，根据表联合主键pid_key_unique_id判断
func (m *JxRedisStorage) SaveKeyByFamilyNode(familyNode *Node) (bool, error) {
 isExist, err := m.IsExistByPidAndKey(RootNodeKeyId, familyNode.Key)
 if err != nil {
 return false, err
 }
 if isExist {
 return m.UpdateKeyByFamily(familyNode)
 }else {
 return m.SetKeyByFamilyNode(familyNode)
 }
}

// IsExistByPidAndKey 根据联合主键来校验数据存在性
func (m *JxRedisStorage) IsExistByPidAndKey(parentId int, key string) (bool, error) {
 if e := global.JX_DB.Model(new(JxRedisStorage)).Where("parent_key_id = ? and key = ?", parentId, key).First(&JxRedisStorage{}).Error; e != nil {
 return false, e
 }
 return true, nil
}
```