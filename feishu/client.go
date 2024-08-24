package feishu

import (
	"context"
	"errors"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkbitable "github.com/larksuite/oapi-sdk-go/v3/service/bitable/v1"
)

const Type = "open_id"

// BitTableClient https://open.feishu.cn/document/server-docs/docs/bitable-v1/app-table/patch
// 无访问权限(1254302 Permission denied.), 常由表格开启了高级权限造成, 如果是用应用请求的话，目前有两种方法对应用赋予高级权限，
// 第一种方法为在表格中添加应用为协作者并将应用设置为管理员，
// 第二种方法为在一个用户群中将应用添加为机器人， 并在高级权限的角色中添加该用户群，从而赋予对应的权限。
type BitTableClient struct {
	appToken string
	tableId  string
	client   *lark.Client
}

// NewBitTableClient 创建多维表格客户端
// appId:
// appSecret:
// appToken: 多维表格 token
// tableId：多维表格 id
func NewBitTableClient(appId, appSecret, appToken, tableId string) *BitTableClient {
	return &BitTableClient{
		client:   lark.NewClient(appId, appSecret),
		appToken: appToken,
		tableId:  tableId,
	}
}

func (b *BitTableClient) Create(ctx context.Context, data map[string]interface{}) (string, error) {
	req := larkbitable.NewCreateAppTableRecordReqBuilder().
		AppToken(b.appToken).
		TableId(b.tableId).
		UserIdType(Type).
		AppTableRecord(larkbitable.NewAppTableRecordBuilder().
			Fields(data).
			Build()).
		Build()

	resp, err := b.client.Bitable.AppTableRecord.Create(ctx, req)
	if err != nil {
		return "", err
	}

	if !resp.Success() {
		return "", errors.New(resp.Error())
	}

	return *resp.Data.Record.RecordId, nil
}

func (b *BitTableClient) Update(ctx context.Context, recordId string, data map[string]interface{}) error {
	updateAppTableRecordReq := larkbitable.NewUpdateAppTableRecordReqBuilder().
		AppToken(b.appToken).
		TableId(b.tableId).
		UserIdType(Type).
		RecordId(recordId).
		AppTableRecord(larkbitable.NewAppTableRecordBuilder().
			Fields(data).
			Build()).
		Build()

	updateResp, err := b.client.Bitable.AppTableRecord.Update(ctx, updateAppTableRecordReq)
	if err != nil {
		return err
	}

	if !updateResp.Success() {
		return errors.New(updateResp.Error())
	}

	return nil
}

func (b *BitTableClient) List(ctx context.Context) (*larkbitable.ListAppTableRecordRespData, error) {
	listAppTableRecordReq := larkbitable.NewListAppTableRecordReqBuilder().
		AppToken(b.appToken).
		TableId(b.tableId).
		UserIdType(Type).
		Build()

	listAppTableRecordResp, err := b.client.Bitable.AppTableRecord.List(ctx, listAppTableRecordReq)
	if err != nil {
		return nil, err
	}

	if !listAppTableRecordResp.Success() {
		return nil, errors.New(listAppTableRecordResp.Error())
	}

	return listAppTableRecordResp.Data, nil
}

func (b *BitTableClient) ListByFilter(ctx context.Context, filter string) (*larkbitable.ListAppTableRecordRespData, error) {
	listAppTableRecordReq := larkbitable.NewListAppTableRecordReqBuilder().
		AppToken(b.appToken).
		TableId(b.tableId).
		Filter(filter).
		UserIdType(Type).
		Build()

	listAppTableRecordResp, err := b.client.Bitable.AppTableRecord.List(ctx, listAppTableRecordReq)
	if err != nil {
		return nil, err
	}

	if !listAppTableRecordResp.Success() {
		return nil, errors.New(listAppTableRecordResp.Error())
	}

	return listAppTableRecordResp.Data, nil
}

func (b *BitTableClient) Get(ctx context.Context, recordId string) (*larkbitable.GetAppTableRecordRespData, error) {
	getAppTableRecordReq := larkbitable.NewGetAppTableRecordReqBuilder().
		AppToken(b.appToken).
		TableId(b.tableId).
		UserIdType(Type).
		RecordId(recordId).
		Build()

	getAppTableRecordResp, err := b.client.Bitable.AppTableRecord.Get(ctx, getAppTableRecordReq)
	if err != nil {
		return nil, err
	}

	if !getAppTableRecordResp.Success() {
		return nil, errors.New(getAppTableRecordResp.Error())
	}

	return getAppTableRecordResp.Data, nil
}
