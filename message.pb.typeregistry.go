// Code generated by "type-registry -in ./message.pb.go -typenameprefix GameMessage_"; DO NOT EDIT.

package message

import "reflect"

var typeregistry = make(map[string]reflect.Type)

func init() {
	typeregistry["CreatePlayerRequest"] = reflect.TypeOf(GameMessage_CreatePlayerRequest{})
	typeregistry["CreatePlayerResponse"] = reflect.TypeOf(GameMessage_CreatePlayerResponse{})
	typeregistry["DataRequest"] = reflect.TypeOf(GameMessage_DataRequest{})
	typeregistry["DataResponse"] = reflect.TypeOf(GameMessage_DataResponse{})
	typeregistry["DeathNotification"] = reflect.TypeOf(GameMessage_DeathNotification{})
	typeregistry["DropNotification"] = reflect.TypeOf(GameMessage_DropNotification{})
	typeregistry["DropRequest"] = reflect.TypeOf(GameMessage_DropRequest{})
	typeregistry["DropResponse"] = reflect.TypeOf(GameMessage_DropResponse{})
	typeregistry["EnterRoomNotification"] = reflect.TypeOf(GameMessage_EnterRoomNotification{})
	typeregistry["EquipRequest"] = reflect.TypeOf(GameMessage_EquipRequest{})
	typeregistry["EquipResponse"] = reflect.TypeOf(GameMessage_EquipResponse{})
	typeregistry["ErrorRequest"] = reflect.TypeOf(GameMessage_ErrorRequest{})
	typeregistry["ErrorResponse"] = reflect.TypeOf(GameMessage_ErrorResponse{})
	typeregistry["ExitInfo"] = reflect.TypeOf(ExitInfo{})
	typeregistry["ExitsRequest"] = reflect.TypeOf(GameMessage_ExitsRequest{})
	typeregistry["ExitsResponse"] = reflect.TypeOf(GameMessage_ExitsResponse{})
	typeregistry["GameMessage"] = reflect.TypeOf(GameMessage{})
	typeregistry["GameMessage_CreatePlayerRequest"] = reflect.TypeOf(GameMessage_CreatePlayerRequest{})
	typeregistry["GameMessage_CreatePlayerResponse"] = reflect.TypeOf(GameMessage_CreatePlayerResponse{})
	typeregistry["GameMessage_DataRequest"] = reflect.TypeOf(GameMessage_DataRequest{})
	typeregistry["GameMessage_DataResponse"] = reflect.TypeOf(GameMessage_DataResponse{})
	typeregistry["GameMessage_DeathNotification"] = reflect.TypeOf(GameMessage_DeathNotification{})
	typeregistry["GameMessage_DropNotification"] = reflect.TypeOf(GameMessage_DropNotification{})
	typeregistry["GameMessage_DropRequest"] = reflect.TypeOf(GameMessage_DropRequest{})
	typeregistry["GameMessage_DropResponse"] = reflect.TypeOf(GameMessage_DropResponse{})
	typeregistry["GameMessage_EnterRoomNotification"] = reflect.TypeOf(GameMessage_EnterRoomNotification{})
	typeregistry["GameMessage_EquipRequest"] = reflect.TypeOf(GameMessage_EquipRequest{})
	typeregistry["GameMessage_EquipResponse"] = reflect.TypeOf(GameMessage_EquipResponse{})
	typeregistry["GameMessage_ErrorRequest"] = reflect.TypeOf(GameMessage_ErrorRequest{})
	typeregistry["GameMessage_ErrorResponse"] = reflect.TypeOf(GameMessage_ErrorResponse{})
	typeregistry["GameMessage_ExitsRequest"] = reflect.TypeOf(GameMessage_ExitsRequest{})
	typeregistry["GameMessage_ExitsResponse"] = reflect.TypeOf(GameMessage_ExitsResponse{})
	typeregistry["GameMessage_GetNotification"] = reflect.TypeOf(GameMessage_GetNotification{})
	typeregistry["GameMessage_GetRequest"] = reflect.TypeOf(GameMessage_GetRequest{})
	typeregistry["GameMessage_GetResponse"] = reflect.TypeOf(GameMessage_GetResponse{})
	typeregistry["GameMessage_InventoryRequest"] = reflect.TypeOf(GameMessage_InventoryRequest{})
	typeregistry["GameMessage_InventoryResponse"] = reflect.TypeOf(GameMessage_InventoryResponse{})
	typeregistry["GameMessage_KillRequest"] = reflect.TypeOf(GameMessage_KillRequest{})
	typeregistry["GameMessage_KillResponse"] = reflect.TypeOf(GameMessage_KillResponse{})
	typeregistry["GameMessage_LeaveRoomNotification"] = reflect.TypeOf(GameMessage_LeaveRoomNotification{})
	typeregistry["GameMessage_LoadRequest"] = reflect.TypeOf(GameMessage_LoadRequest{})
	typeregistry["GameMessage_LoadResponse"] = reflect.TypeOf(GameMessage_LoadResponse{})
	typeregistry["GameMessage_LoginRequest"] = reflect.TypeOf(GameMessage_LoginRequest{})
	typeregistry["GameMessage_LoginResponse"] = reflect.TypeOf(GameMessage_LoginResponse{})
	typeregistry["GameMessage_LogoutNotification"] = reflect.TypeOf(GameMessage_LogoutNotification{})
	typeregistry["GameMessage_LogoutRequest"] = reflect.TypeOf(GameMessage_LogoutRequest{})
	typeregistry["GameMessage_LogoutResponse"] = reflect.TypeOf(GameMessage_LogoutResponse{})
	typeregistry["GameMessage_LookNotification"] = reflect.TypeOf(GameMessage_LookNotification{})
	typeregistry["GameMessage_LookRequest"] = reflect.TypeOf(GameMessage_LookRequest{})
	typeregistry["GameMessage_LookResponse"] = reflect.TypeOf(GameMessage_LookResponse{})
	typeregistry["GameMessage_MoveRequest"] = reflect.TypeOf(GameMessage_MoveRequest{})
	typeregistry["GameMessage_MoveResponse"] = reflect.TypeOf(GameMessage_MoveResponse{})
	typeregistry["GameMessage_Ping"] = reflect.TypeOf(GameMessage_Ping{})
	typeregistry["GameMessage_Pong"] = reflect.TypeOf(GameMessage_Pong{})
	typeregistry["GameMessage_RecallRequest"] = reflect.TypeOf(GameMessage_RecallRequest{})
	typeregistry["GameMessage_RecallResponse"] = reflect.TypeOf(GameMessage_RecallResponse{})
	typeregistry["GameMessage_RoomStatusRequest"] = reflect.TypeOf(GameMessage_RoomStatusRequest{})
	typeregistry["GameMessage_RoomStatusResponse"] = reflect.TypeOf(GameMessage_RoomStatusResponse{})
	typeregistry["GameMessage_SayNotification"] = reflect.TypeOf(GameMessage_SayNotification{})
	typeregistry["GameMessage_SayRequest"] = reflect.TypeOf(GameMessage_SayRequest{})
	typeregistry["GameMessage_SayResponse"] = reflect.TypeOf(GameMessage_SayResponse{})
	typeregistry["GameMessage_ShowEquipmentRequest"] = reflect.TypeOf(GameMessage_ShowEquipmentRequest{})
	typeregistry["GameMessage_ShowEquipmentResponse"] = reflect.TypeOf(GameMessage_ShowEquipmentResponse{})
	typeregistry["GameMessage_StatRequest"] = reflect.TypeOf(GameMessage_StatRequest{})
	typeregistry["GameMessage_StatResponse"] = reflect.TypeOf(GameMessage_StatResponse{})
	typeregistry["GameMessage_TellAllNotification"] = reflect.TypeOf(GameMessage_TellAllNotification{})
	typeregistry["GameMessage_TellAllRequest"] = reflect.TypeOf(GameMessage_TellAllRequest{})
	typeregistry["GameMessage_TellAllResponse"] = reflect.TypeOf(GameMessage_TellAllResponse{})
	typeregistry["GameMessage_TellNotification"] = reflect.TypeOf(GameMessage_TellNotification{})
	typeregistry["GameMessage_TellRequest"] = reflect.TypeOf(GameMessage_TellRequest{})
	typeregistry["GameMessage_TellResponse"] = reflect.TypeOf(GameMessage_TellResponse{})
	typeregistry["GameMessage_ViolenceNotification"] = reflect.TypeOf(GameMessage_ViolenceNotification{})
	typeregistry["GameMessage_WearRequest"] = reflect.TypeOf(GameMessage_WearRequest{})
	typeregistry["GameMessage_WearResponse"] = reflect.TypeOf(GameMessage_WearResponse{})
	typeregistry["GameMessage_WhoRequest"] = reflect.TypeOf(GameMessage_WhoRequest{})
	typeregistry["GameMessage_WhoResponse"] = reflect.TypeOf(GameMessage_WhoResponse{})
	typeregistry["GetNotification"] = reflect.TypeOf(GameMessage_GetNotification{})
	typeregistry["GetRequest"] = reflect.TypeOf(GameMessage_GetRequest{})
	typeregistry["GetResponse"] = reflect.TypeOf(GameMessage_GetResponse{})
	typeregistry["InventoryRequest"] = reflect.TypeOf(GameMessage_InventoryRequest{})
	typeregistry["InventoryResponse"] = reflect.TypeOf(GameMessage_InventoryResponse{})
	typeregistry["InventoryResponse_InventoryItem"] = reflect.TypeOf(InventoryResponse_InventoryItem{})
	typeregistry["KillRequest"] = reflect.TypeOf(GameMessage_KillRequest{})
	typeregistry["KillResponse"] = reflect.TypeOf(GameMessage_KillResponse{})
	typeregistry["LeaveRoomNotification"] = reflect.TypeOf(GameMessage_LeaveRoomNotification{})
	typeregistry["LoadRequest"] = reflect.TypeOf(GameMessage_LoadRequest{})
	typeregistry["LoadResponse"] = reflect.TypeOf(GameMessage_LoadResponse{})
	typeregistry["LoginRequest"] = reflect.TypeOf(GameMessage_LoginRequest{})
	typeregistry["LoginResponse"] = reflect.TypeOf(GameMessage_LoginResponse{})
	typeregistry["LogoutNotification"] = reflect.TypeOf(GameMessage_LogoutNotification{})
	typeregistry["LogoutRequest"] = reflect.TypeOf(GameMessage_LogoutRequest{})
	typeregistry["LogoutResponse"] = reflect.TypeOf(GameMessage_LogoutResponse{})
	typeregistry["LookNotification"] = reflect.TypeOf(GameMessage_LookNotification{})
	typeregistry["LookRequest"] = reflect.TypeOf(GameMessage_LookRequest{})
	typeregistry["LookResponse"] = reflect.TypeOf(GameMessage_LookResponse{})
	typeregistry["MoveRequest"] = reflect.TypeOf(GameMessage_MoveRequest{})
	typeregistry["MoveResponse"] = reflect.TypeOf(GameMessage_MoveResponse{})
	typeregistry["Ping"] = reflect.TypeOf(GameMessage_Ping{})
	typeregistry["Pong"] = reflect.TypeOf(GameMessage_Pong{})
	typeregistry["RecallRequest"] = reflect.TypeOf(GameMessage_RecallRequest{})
	typeregistry["RecallResponse"] = reflect.TypeOf(GameMessage_RecallResponse{})
	typeregistry["RoomDescription"] = reflect.TypeOf(RoomDescription{})
	typeregistry["RoomStatusRequest"] = reflect.TypeOf(GameMessage_RoomStatusRequest{})
	typeregistry["RoomStatusResponse"] = reflect.TypeOf(GameMessage_RoomStatusResponse{})
	typeregistry["RoomStatusResponse_DirectionInfo"] = reflect.TypeOf(RoomStatusResponse_DirectionInfo{})
	typeregistry["RoomStatusResponse_InventoryInfo"] = reflect.TypeOf(RoomStatusResponse_InventoryInfo{})
	typeregistry["RoomStatusResponse_MobInfo"] = reflect.TypeOf(RoomStatusResponse_MobInfo{})
	typeregistry["RoomStatusResponse_PlayerInfo"] = reflect.TypeOf(RoomStatusResponse_PlayerInfo{})
	typeregistry["SayNotification"] = reflect.TypeOf(GameMessage_SayNotification{})
	typeregistry["SayRequest"] = reflect.TypeOf(GameMessage_SayRequest{})
	typeregistry["SayResponse"] = reflect.TypeOf(GameMessage_SayResponse{})
	typeregistry["ShowEquipmentRequest"] = reflect.TypeOf(GameMessage_ShowEquipmentRequest{})
	typeregistry["ShowEquipmentResponse"] = reflect.TypeOf(GameMessage_ShowEquipmentResponse{})
	typeregistry["ShowEquipmentResponse_EquipmentInfo"] = reflect.TypeOf(ShowEquipmentResponse_EquipmentInfo{})
	typeregistry["StatRequest"] = reflect.TypeOf(GameMessage_StatRequest{})
	typeregistry["StatResponse"] = reflect.TypeOf(GameMessage_StatResponse{})
	typeregistry["TellAllNotification"] = reflect.TypeOf(GameMessage_TellAllNotification{})
	typeregistry["TellAllRequest"] = reflect.TypeOf(GameMessage_TellAllRequest{})
	typeregistry["TellAllResponse"] = reflect.TypeOf(GameMessage_TellAllResponse{})
	typeregistry["TellNotification"] = reflect.TypeOf(GameMessage_TellNotification{})
	typeregistry["TellRequest"] = reflect.TypeOf(GameMessage_TellRequest{})
	typeregistry["TellResponse"] = reflect.TypeOf(GameMessage_TellResponse{})
	typeregistry["UnimplementedMudCommServer"] = reflect.TypeOf(UnimplementedMudCommServer{})
	typeregistry["ViolenceNotification"] = reflect.TypeOf(GameMessage_ViolenceNotification{})
	typeregistry["WearRequest"] = reflect.TypeOf(GameMessage_WearRequest{})
	typeregistry["WearResponse"] = reflect.TypeOf(GameMessage_WearResponse{})
	typeregistry["WhoRequest"] = reflect.TypeOf(GameMessage_WhoRequest{})
	typeregistry["WhoResponse"] = reflect.TypeOf(GameMessage_WhoResponse{})
	typeregistry["WhoResponse_PlayerInfo"] = reflect.TypeOf(WhoResponse_PlayerInfo{})
}
