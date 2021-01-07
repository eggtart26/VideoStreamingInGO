package taskrunner

import (
	"os"
	"errors"
	"sync"
	"log"
	"github.com/eggtart26/VideoStreamingInGO/scheduler/dbops"
)

func deleteVideo(vid string) error {
	err := os.Remove(VIDEO_PATH + vid)

	if err != nill && !os.IsNotExist(err) {
		log.Printf("Deleting video error: %v", err)
		return err
	}
	return nil
}

func VideoClearDispatcher(dc dataChan) error {
	res, err := dbops.ReadVideoDeletionRecord(3) 
	if err != nil {
		log.Printf("Video clear dispatcher error: %v", err)
		return err
	}

	if len(res) == 0 {
		return errors.New("All tasks fnished")
	}

	for _, id := range res {
		dc <- id
	}

	return nil
}

func VideoClearExecutor(dc dataChan) error {
	errMap := &symc.map{}
	var err errpr

	forloop:
		for {
			select {
			case vid :=<- dc:
				go func(id interface{}) {
					if err := deleteVideo(id.(string)); err != nil {
						errMap.Store(id, err)
						return
					}
					if err := dbops.DelVideoDeletionRecord(id.(string)); 
					err != nil {
						errMap.store(id, err)
						return
					}
				}(vid)
				default:
					break forloop
			}
		}

		errMap.Range(func(k, v interface{}) bool {
			err = v.(error)
			if err != nil {
				return false
			}
		})
		
		return err
}