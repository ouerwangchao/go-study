package interview

func verifyPostorder(postorder []int) bool {
	if len(postorder) <= 1{
		return true
	}
	res := check(postorder,0,len(postorder)-1)
	return res
}

func check(postorder []int,start,end int) bool{
	//后序遍历，最后一个值为根节点
	if start >= end{
		return true
	}
	root := postorder[end]
	i := start
	//定位根节点的左右子节点分界
	for postorder[i]< root{
		i++
	}
	//判断postorder数组中i以后的元素是否有小于root的，有则返回false
	j := i
	for j<end{
		if postorder[j] < root{
			return false
		}
		j++
	}

	return check(postorder,start,i-1) && check(postorder,i,end-1)
}