package item

import (
	"bufio"
	"encoding/json"
	"fmt"
	"time"

	"github.com/AswinKakarot/redis-populator/pkg/config"
	"github.com/AswinKakarot/redis-populator/pkg/redisutil"
	"github.com/AswinKakarot/redis-populator/pkg/util"
)

var (
	itemStore ItemStore
)

func NewItemStore(ls *util.LocalStore, r redisutil.Redis) *ItemStore {
	itemStore.Items = make([]Item, 0)
	itemStore.LocalStore = ls
	itemStore.Redis = r
	return &itemStore
}

type ItemStore struct {
	Items []Item

	LocalStore *util.LocalStore
	Redis      redisutil.Redis
	rSuccessCount,
	rFailCount,
	wSuccessCount,
	wFailCount,
	MaxKeyLen,
	MaxValueLen,
	MinKeyLen,
	MinValueLen int
}

func (it *ItemStore) Write() error {
	appCfg := config.GetAppConfig()
	for i := 0; i < appCfg.WriteDuration; i++ {
		item := Item{
			Key:   util.RandomKey(it.MinKeyLen, it.MaxKeyLen),
			Value: util.RandomValue(it.MinValueLen, it.MaxValueLen),
		}
		if err := it.Redis.Write(item.Key, item.Value); err != nil {
			it.wFailCount++
			return err
		}
		out, err := json.Marshal(item)
		if err != nil {
			it.wFailCount++
			return err
		}
		if it.LocalStore != nil {
			if _, err := fmt.Fprintln(it.LocalStore.Fptr, string(out)); err != nil {
				it.wFailCount++
				return err
			}
		}
		fmt.Println("Written Data:", string(out))
		it.Items = append(it.Items, item)
		time.Sleep(time.Second * time.Duration(appCfg.WriteInterval))
		it.wSuccessCount++
	}
	return nil
}

func (it *ItemStore) Read(readFromLocal bool) (err error) {
	appCfg := config.GetAppConfig()
	var pos int
	var scanner *bufio.Scanner
	if it.LocalStore != nil {
		scanner = bufio.NewScanner(it.LocalStore.Fptr)
	}
	for i := 0; i < appCfg.ReadDuration; i++ {
		time.Sleep(time.Second * time.Duration(appCfg.ReadInterval))
		if readFromLocal {
			if scanner == nil {
				it.rFailCount++
				return fmt.Errorf("nothing to read from")
			}
			if !scanner.Scan() {
				it.rFailCount++
				return fmt.Errorf("end of file")
			}
			item := Item{}
			if err = json.Unmarshal(scanner.Bytes(), &item); err != nil {
				return
			}
			it.Items = append(it.Items, item)
			item.Value, err = it.Redis.Read(item.Key)
			if err != nil {
				it.rFailCount++
				return
			}
		} else if len(it.Items) == 0 {
			it.rFailCount++
			return fmt.Errorf("nothing to read from")
		} else {
			if len(it.Items) <= i {
				pos = 0
			}
			item := it.Items[pos]
			item.Value, err = it.Redis.Read(item.Key)
			if err != nil {
				it.rFailCount++
				return
			}
			fmt.Println("Read Value:", item.Value)
			pos++
		}
		it.rSuccessCount++
	}
	return
}

func (it *ItemStore) PrintWriteStatus() {
	fmt.Println("Write Status")
	fmt.Println("============")
	fmt.Printf("Successful: %d\n", it.wSuccessCount)
	fmt.Printf("Failure: %d\n", it.wFailCount)
	fmt.Println()
}

func (it *ItemStore) PrintReadStatus() {
	fmt.Println("Read Status")
	fmt.Println("===========")
	fmt.Printf("Successful: %d\n", it.rSuccessCount)
	fmt.Printf("Failure: %d\n", it.rFailCount)
	fmt.Println()
}
