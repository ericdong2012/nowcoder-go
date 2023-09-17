# 它是一种基于比较的算法，其中每对相邻元素进行比较，如果元素不合适，元素将进行交换。


def bubblesort(list):
    # Swap the elements to arrange in order
    for iter_num in range(len(list) - 1, 0, -1):
        for idx in range(iter_num):
            if list[idx] > list[idx + 1]:
                temp = list[idx]
                list[idx] = list[idx + 1]
                list[idx + 1] = temp


# func  bubbleSort(list []int) []int {
# 	for i:=0 ; i < len(list); i++ {
# 		for j:= i+1 ; j < len(list ) ; j++ {
# 			if list[i] > list[j] {
# 			 	list[i], list[j] = list[j], list[i]
# 			 }
# 		}
# 	}
# 	return list
# }

# func  selectSort(list []int) []int {
# 	for i:=0; i < len(list); i ++ {
# 		minI := i
# 		for j:=1; j < len(list ); j++ {
# 			if list[minI] > list[j] {
# 				minI = j
# 			}
# 		}
# 		list[i], list[minI] = list[minI], list[i]
# 	}
# 	return list
# }


list = [19, 2, 31, 45, 6, 11, 121, 27]
bubblesort(list)
print(list)

print(11111)
print(2222)



# [2, 6, 11, 19, 27, 31, 45, 121]
