type Flight struct {
    Id string
    DepTime time.Time
}

type Dependency struct {
    ParentId string
    ChildId string
    Buffer time.Duration
}

type DelayedFlight struct {
    Id string
    Delay time.Duration
}

func FindDelayedFlights(flights []Flight, dependencies []Dependency, delayedFlightId string, delayTime time.Duration) []DelayedFlight {  
    var flightMap map[string]*Flight = make(map[string]*Flight)
    var dependencyMap map[string][]Dependency = make(map[string][]Dependency)
    var delayMap = make(map[string]*DelayedFlight)
    
    for i :=0; i< len(flights); i++ {
        flightMap[flights[i].Id] = &flights[i]
    }

    for i :=0; i< len(dependencies); i++ {
        dep := dependencies[i]
        dependencyMap[dep.ParentId] = append(dependencyMap[dep.ParentId], dep)
    }
    
    delayMap[delayedFlightId] = &DelayedFlight{Id: delayedFlightId, Delay: delayTime}

    var queue []DelayedFlight = []DelayedFlight{}
    finalDelayQueue := make([]DelayedFlight, 0)
    
    queue = append(queue, DelayedFlight{Id: delayedFlightId, Delay: delayTime})

    for len(queue)>0 {
        var newQueue []DelayedFlight
        for i := 0; i<len(queue); i++ {
            currFlight := queue[i]
            currDependencies, ok := dependencyMap[currFlight.Id]
            if ok && len(currDependencies)>0 {
                for j :=0; j< len(currDependencies); j++ {
                    currDep := currDependencies[j]
                    delay :=  currFlight.Delay + currDep.Buffer - flightMap[currDep.ChildId].DepTime.Sub(flightMap[currFlight.Id].DepTime)
                    if ( delay > 0) {
                        prevDelay, ok := delayMap[currDep.ChildId]
                        if ok {
                            if  prevDelay.Delay < delay{
                                prevDelay.Delay = delay
                            newQueue = append(newQueue, DelayedFlight{Id: currDep.ChildId, Delay: delay})
                            }
                        }else {
                            delayMap[currDep.ChildId] = &DelayedFlight{Id: currDep.ChildId, Delay: delay}
                            newQueue = append(newQueue, DelayedFlight{Id: currDep.ChildId, Delay: delay})
                        }
                    }
                }
            }
        }

        queue = newQueue
    }

    for _, val := range delayMap {
        finalDelayQueue = append(finalDelayQueue, *val)
    }

    return finalDelayQueue
}