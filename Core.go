package main

func check(e error) {
    if e != nil{
        panic(e)
    }
}

func find(slice []int, val int)(bool){
    for _, item := range slice{
        if item == val{
            return true
        }
    }
    return false
}