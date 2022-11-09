pub fn remove_element(nms: &mut Vec<i32>, val: i32) -> i32 {
    let mut slow = 0;
    for i in 0..nms.len() {
        if nms[i] != val {
            nms[slow] = nms[i];
            slow += 1;
        };
    }
    return slow;
}
