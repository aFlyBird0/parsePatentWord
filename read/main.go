package read

func Get(filename string) (parasCombine Paras){
	parasOriginal := Read(filename)
	parasCombine = parasOriginal.Combine()
	return
}
