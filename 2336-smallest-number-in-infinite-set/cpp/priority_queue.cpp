class SmallestInfiniteSet {
public:
    std::unordered_set<int> inPQ;
    std::priority_queue<int, std::vector<int>, std::greater<int>> pq;
    int next;

    SmallestInfiniteSet() {
        this->next = 1;
    }
    
    int popSmallest() {
        int small;
        if (pq.size() > 0 && pq.top() < next) {
            small = pq.top();
            inPQ.erase(small);
            pq.pop();
        } else {
            small = next++;
        }

        return small;
    }
    
    void addBack(int num) {
        if(this->next <= num) return;
        if(this->inPQ.find(num) != this->inPQ.end()) return;
        
        this->pq.push(num);
        this->inPQ.insert(num);
    }
};
/* 
    pq -> []
    next -> 1

    addBack(num) -> 
        if next >= num no se hace nada
        pq.push(num);



    popSmalltest() -> 
        small = next;
        if pq.front() < next 
            small = pq.front();
            pq.pop();
        else
            small = next++;

        return small
*/

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * SmallestInfiniteSet* obj = new SmallestInfiniteSet();
 * int param_1 = obj->popSmallest();
 * obj->addBack(num);
 */