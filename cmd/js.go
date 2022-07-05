package main

type Array []interface{}

func (arr *Array) ForEach(fn func(interface{})) {

	for i := 0; i < len(*arr); i++ {

		fn((*arr)[i])

	}

}
