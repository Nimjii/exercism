package twofer

func ShareWith(name string) string {
    var call string
    
    if name == "" {
        call = "you"
    } else {
        call = name
    }
    
	return "One for " + call + ", one for me."
}
