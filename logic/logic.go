package logic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

type block struct {
	Cid          string `json:"cid"`
	Height       int64  `json:"height"`
	Timestamp    int64  `json:"timestamp"`
	Size         int64  `json:"size"`
	WinCount     int64  `json:"winCount"`
	Reward       string `json:"reward"`
	Penalty      string `json:"penalty"`
	MessageCount int64  `json:"messageCount"`
}
type data struct {
	TotalCount int64   `json:"totalCount"`
	Blocks     []block `json:"blocks"`
}
type R struct {
	C      string `json:"code"`
	Result struct {
		Datetime []string `json:"datetime"`
		Powers   []string `json:"powers"`
	} `json:"data"`
}
type PowerIn7D struct {
	Datetime []string `json:"datetime"`
	Powers   []string `json:"powers"`
}

type GasRsp struct {
	C      string `json:"code"`
	Result struct {
		BaseFee struct {
			unit    string    `json:"unit"`
			FeeList []float64 `json:"feeList"`
		} `json:"baseFee"`

		GasFee64 struct {
			unit    string    `json:"unit"`
			FeeList []float64 `json:"feeList"`
		} `json:"gasFee64"`
		GasFee32 struct {
			unit    string    `json:"unit"`
			FeeList []float64 `json:"feeList"`
		} `json:"gasFee32"`

		Height   []uint64 `json:"height"`
		TimeList []string `json:"timeList"`
	} `json:"data"`
}

type Russ struct {
	Datetime string `json:"datetime"`
	Powers   string `json:"powers"`
}
type Gas struct {
	BaseFee  float64 `json:"baseFee"`
	GasFee32 float64 `json:"gasFee32"`
	GasFee64 float64 `json:"gasFee64"`
	Height   uint64  `json:"height"`
	Time     string  `json:"timeList"`
}
type PowerRepo struct {
	C      string `json:"code"`
	Result struct {
		Data []struct {
			HeightTimeStr         string `json:"heightTimeStr"`
			QualityPowerGrowthStr string `json:"qualityPowerGrowthStr"`
		} `json:"powers"`
	} `json:"data"`
}

type PowerRes struct {
	HeightTimeStr         string `json:"heightTimeStr"`
	QualityPowerGrowthStr string `json:"qualityPowerGrowthStr"`
}
type BlockReward struct {
	Result struct {
		Basic struct {
			BlockReward string `json:"rewards"`
		} `json:"basic"`
	} `json:"result"`
}

type InfoByAddress struct {
	Data struct {
		Miner               string  `json:"miner"`
		QualityPower        int64   `json:"qualitypower"`
		QualityPowerStr     string  `json:"qualitypowerstr"`
		QualityPowerPercent float64 `json:"qualitypowerpercent"`
		RawPower            int64   `json:"rawpower"`
		RawPowerStr         string  `json:"rawpowerstr"`
		RawPowerPercent     string  `json:"rawpowerpercent"`
		TotalQualityPower   int64   `json:"totalqualitypower"`
		TotalRawPower       int64   `json:"totalrawpower"`
		TotalRawPowerStr    string  `json:"totalrawpowerstr"`
		Blocks              int64   `json:"blocks"`
		WinCount            int64   `json:"wincount"`
		BlockReward         string  `json:"blockrewad"`
		Owner               string  `json:"owner"`
		Worker              string  `json:"worker"`
		Tag                 string  `json:"tag"`
		IsVerified          int64   `json:"isverified"`
		PeerId              string  `json:"peerid"`
		PowerRank           int64   `json:"powerrank"`
		Local               struct {
			IP       string `json:"ip"`
			Location string `json:"location"`
		} `json:"local"`
		Balance struct {
			Balance          float64 `json:"balance"`
			BalanceStr       string  `json:"balancestr"`
			Available        float64 `json:"available"`
			AvailableStr     string  `json:"availablestr"`
			SectorsPledge    float64 `json:"sectorspledge"`
			SectorsPledgeStr string  `json:"sectorspledgestr"`
			LockedFunds      float64 `json:"lockedfunds"`
			LockedFundsStr   string  `json:"lockedfundsstr"`
			FeeDebtStr       string  `json:"feedebtstr"`
		} `json:"balance"`
		Sector struct {
			SectorSize    int64  `json:"sectorsize"`
			SectorSizeStr string `json:"sectorsizestr"`
			SectorCount   int64  `json:"sectorcount"`
			ActiveCount   int64  `json:"activecount"`
			FaultCount    int64  `json:"faultcount"`
			RecoveryCount int64  `json:"recoverycount"`
		} `json:"sectors"`
	} `json:"data"`
}

type base struct {
	Data struct {
		AddPowerIn32g               string  `json:"add_power_in_32g"`
	} `json:"data"`
}



type MiningStatus struct {
	RawBytePowerGrowth     string  `json:"rawBytePowerGrowth"`
	QualityAdjPowerGrowth  string  `json:"qualityAdjPowerGrowth"`
	AwBytePowerDelta       string  `json:"rawBytePowerDelta"`
	QualityAdjPowerDelta   string  `json:"qualityAdjPowerDelta"`
	BlocksMined            int64   `json:"blocksMined"`
	WeightedBlocksMined    int64   `json:"weightedBlocksMined"`
	TotalRewards           string  `json:"totalRewards"`
	NetworkTotalRewards    string  `json:"networkTotalRewards"`
	EquivalentMiners       float64 `json:"equivalentMiners"`
	RewardPerByte          float64 `json:"rewardPerByte"`
	WindowedPoStFeePerByte float64 `json:"windowedPoStFeePerByte"`
	LuckyValue             float64 `json:"luckyValue"`
	DurationPercentage     int     `json:"durationPercentage"`
}
//gas_in_64g: "0.00014725490533037029"
type BaseInfo struct {
	Data struct {
		TotalPower       float64 `json:"totalPower"`
		PledgeCollateral float64 `json:"pledgeCollateral"`
		PowerIn24H       string  `json:"PowerIn24H"`
		NewlyFilIn24h    float64 `json:"newlyFilIn24h"`
		BlockRewardIn24h float64 `json:"blockRewardIn24h"`
		CurrentPledgeCollateral float64 `json:"currentPledgeCollateral"`
		NewlyPowerCostIn32GB float64 `json:"newlyPowerCostIn32GB"`
		GasIn32G string `json:"gasIn32g"`
		GasIn64G string `json:"gasIn64g"`
	} `json:"data"`
}

type Base struct {
	Result struct {
		Data struct {
			NewlyPowerCostIn32GB string `json:"add_power_in_32g"`
			GasIn32G string `json:"gas_in_32g"`
			GasIn64G string `json:"gas_in_64g"`
		} `json:"data"`
	} `json:"result"`
}
//	//	"id":1,
//	//	"jsonrpc":"2.0",
//	//	"method":"filscan.ActorById",
//	//	"params":["f0469055"]
type Body struct {
	ID      int      `json:"id"`
	JsonRpc string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
}

func MinerByPeerId(minerId string) ([]block, error) {
	MinerUrl := fmt.Sprintf("https://filfox.info/api/v1/address/%s/blocks?pageSize=30&page=0", minerId)
	Database := new(data)
	params := url.Values{}
	Url, err := url.Parse(MinerUrl)
	if err != nil {
		return nil, err
	}
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := http.Get(urlPath)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	Data := new(data)
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &Data)
	res := int(Data.TotalCount / 100)
	if res > 0 {
		for i := 0; i < res; i++ {
			params.Set("pageSize", "100")
			pages := fmt.Sprintf("%v", i)
			params.Set("page", pages)
			Url.RawQuery = params.Encode()
			urlPath := Url.String()
			resp, err := http.Get(urlPath)
			if err != nil {
				panic(err)
			}

			Data := new(data)
			body, err := ioutil.ReadAll(resp.Body)
			err = json.Unmarshal(body, &Data)
			for _, block := range Data.Blocks {
				Database.Blocks = append(Database.Blocks, block)
			}
		}
		defer resp.Body.Close()
		return Database.Blocks, err
	}
	return nil, err
}

func GetPowerInByAddress(minerId string) ([]PowerRes, error) {
	Database := new(PowerRepo)
	var res []PowerRes
	MinerUrl := fmt.Sprintf("https://api.filscout.com/api/v1/miners/%s/powerstats", minerId)
	result, err := Post(MinerUrl, `{"statsType":"30d"}"`, "application/json; charset=utf-8")
	if err != nil {
		return res, err
	}
	json.Unmarshal(result, &Database)
	for i := 0; i < len(Database.Result.Data); i++ {
		S := PowerRes{
			HeightTimeStr:         Database.Result.Data[i].HeightTimeStr,
			QualityPowerGrowthStr: Database.Result.Data[i].QualityPowerGrowthStr,
		}
		json.NewEncoder(os.Stdout).Encode(&S)
		res = append(res, S)
	}
	return res, err
}

func GetInfoByAddress(address string) (*InfoByAddress, error) {
	var addre = []string{address}

	Database := new(InfoByAddress)
	BlockRe := new(BlockReward)
	mURL := "https://api.filscan.io:8700/rpc/v1"
	S := Body{
		ID:      1,
		JsonRpc: "2.0",
		Method:  "filscan.ActorById",
		Params:  addre,
	}
	data, err := json.Marshal(&S)
	if err != nil {
		return Database, err
	}
	req, err := http.NewRequest("POST", mURL, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	url := fmt.Sprintf("https://api.filscout.com/api/v1/miners/%s/info", address)
	result, err := Post(url, "", "application/json; charset=utf-8")
	if err != nil {
		return Database, err
	}
	_ = json.Unmarshal(body, &BlockRe)
	_ = json.Unmarshal(result, &Database)
	Database.Data.BlockReward = BlockRe.Result.Basic.BlockReward
	return Database, err
}

func MiningStats(address, type_status string) (*MiningStatus, error) {
	Database := new(MiningStatus)
	Murl := fmt.Sprintf("https://filfox.info/api/v1/address/%s/mining-stats?duration=%s", address,type_status)
	body,err := Get1(Murl)
	if err != nil {
		return Database, err
	}
	json.Unmarshal(body, &Database)
	return Database, err
}

func BaseInfoFun() (*BaseInfo, error) {
	MinerUrl := fmt.Sprintf("https://api.filscout.com/api/v1/network/overview/info")
	Database := new(BaseInfo)
	Base := new(Base)
	params := url.Values{}
	Url, err := url.Parse(MinerUrl)
	if err != nil {
		return nil, err
	}
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := http.Get(urlPath)
	if err != nil {
		if err != nil {
			return nil, err
		}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(body, &Database)
	var addre []string
	mURL := "https://api.filscan.io:8700/rpc/v1"
	S := Body{
		ID:      1,
		JsonRpc: "2.0",
		Method:   "filscan.StatChainInfo",
		Params:  addre,
	}
	data, err := json.Marshal(&S)
	if err != nil {
		return Database, err
	}
	req, err := http.NewRequest("POST", mURL, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp1, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body1, _ := ioutil.ReadAll(resp1.Body)
	_ = json.Unmarshal(body1, &Base)
	score, _ := strconv.ParseFloat(Base.Result.Data.NewlyPowerCostIn32GB, 64)
	fmt.Println(Base.Result.Data.GasIn32G)
	Database.Data.NewlyPowerCostIn32GB = score
	Database.Data.GasIn32G = Base.Result.Data.GasIn32G
	Database.Data.GasIn64G = Base.Result.Data.GasIn64G
	return Database, err
}

func GetPowerIn() ([]Russ, error) {
	Database := new(R)
	var res []Russ
	result, err := Post("https://api.filscout.com/api/v1/tipset/powerchange/month", "", "application/json; charset=utf-8")
	if err != nil {
		return res, err
	}
	json.Unmarshal(result, &Database)
	for i := 0; i < len(Database.Result.Datetime); i++ {
		S := Russ{
			Datetime: Database.Result.Datetime[i],
			Powers:   Database.Result.Powers[i],
		}
		json.NewEncoder(os.Stdout).Encode(&S)
		res = append(res, S)
	}
	return res, err
}

func GetGas(s string) ([]Gas, error) {
	Database := new(GasRsp)
	var res []Gas
	url := fmt.Sprintf("https://api.filscout.com/api/v1/block/feechart/%s", s)
	result, err := Post(url, "", "application/json; charset=utf-8")
	if err != nil {
		return res, err
	}
	json.Unmarshal(result, &Database)
	for i := 0; i < len(Database.Result.Height); i++ {
		S := Gas{
			Height:   Database.Result.Height[i],
			BaseFee:  Database.Result.BaseFee.FeeList[i],
			GasFee32: Database.Result.GasFee32.FeeList[i],
			GasFee64: Database.Result.GasFee64.FeeList[i],
			Time:     Database.Result.TimeList[i],
		}
		json.NewEncoder(os.Stdout).Encode(&S)
		res = append(res, S)
	}
	return res, err
}

// Post 发送POST请求
//url:请求地址		data:POST请求提交的数据		contentType:请求体格式，如：application/json
//content:请求返回的内容
func Post(url string, data string, contentType string) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	req.Header.Add("content-type", contentType)
	if err != nil {
		return []byte(""), err
	}
	defer req.Body.Close()
	client := &http.Client{Timeout: 5 * time.Second}
	resp, error := client.Do(req)
	if error != nil {
		return []byte(""), error
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	return result, err
}

// Get1 发送GET请求
//url:请求地址
//response:请求返回的内容
func Get1(url string) ([]byte, error) {
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	defer resp.Body.Close()
	if err != nil {

		return []byte(""), err
	}
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			return []byte(""), err
		}
	}
	return result.Bytes(), nil
}

