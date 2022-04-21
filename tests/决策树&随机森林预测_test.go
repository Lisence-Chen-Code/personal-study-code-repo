package test

//
//func TestDecisionTree(t *testing.T) {
//	start := time.Now()
//	f, _ := os.Open("train.csv")
//	defer f.Close()
//	t, _ := os.Open("test.csv")
//	defer t.Close()
//	content, _ := ioutil.ReadAll(f)
//	s_content := string(content)
//	lines := strings.Split(s_content, "\n")
//	inputs := make([][]interface{}, 0)
//	targets := make([]string, 0)
//	for _, line := range lines {
//
//		line = strings.TrimRight(line, "\r\n")
//
//		if len(line) == 0 {
//			continue
//		}
//		tup := strings.Split(line, ",")
//		//特征
//		pattern := tup[:len(tup)-1]
//		//结果
//		target := tup[len(tup)-1]
//		X := make([]interface{}, 0)
//		for _, x := range pattern {
//			X = append(X, x)
//		}
//		//特征集
//		inputs = append(inputs, X)
//		//结果集
//		targets = append(targets, target)
//	}
//	tcontent, _ := ioutil.ReadAll(f)
//	t_content := string(tcontent)
//	tlines := strings.Split(t_content, "\n")
//	tinputs := make([][]interface{}, 0)
//	ttargets := make([]string, 0)
//	for _, line := range tlines {
//
//		line = strings.TrimRight(line, "\r\n")
//
//		if len(line) == 0 {
//			continue
//		}
//		tup := strings.Split(line, ",")
//		//特征
//		pattern := tup[:len(tup)-1]
//		//结果
//		target := tup[len(tup)-1]
//		X := make([]interface{}, 0)
//		for _, x := range pattern {
//			X = append(X, x)
//		}
//		//特征集
//		tinputs = append(inputs, X)
//		//结果集
//		ttargets = append(targets, target)
//	}
//	//训练集
//	train_inputs := inputs
//	train_targets := targets
//	//测试集
//	test_inputs := tinputs
//	test_targets := ttargets
//
//	//构建随机森林
//	forest := RF.BuildForest(inputs, targets, 10, 500, len(train_inputs[0])) //10 trees
//
//	test_inputs = train_inputs
//	test_targets = train_targets
//	err_count := 0.0
//	//测试准确率
//	for i := 0; i < len(test_inputs); i++ {
//		output := forest.Predicate(test_inputs[i])
//		expect := test_targets[i]
//		if output != expect {
//			err_count += 1
//		}
//	}
//
//	fmt.Println("success rate:", 1.0-err_count/float64(len(test_inputs)))
//
//	fmt.Println(time.Since(start))
//
//}
