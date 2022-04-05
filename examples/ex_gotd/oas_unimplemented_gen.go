// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"math/big"
	"math/bits"
	"net"
	"net/http"
	"net/netip"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/instrument/syncint64"
	"go.opentelemetry.io/otel/metric/nonrecording"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// No-op definition for keeping imports.
var (
	_ = bytes.NewReader
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = io.Copy
	_ = math.Mod
	_ = big.Rat{}
	_ = bits.LeadingZeros64
	_ = net.IP{}
	_ = http.MethodGet
	_ = netip.Addr{}
	_ = url.URL{}
	_ = regexp.MustCompile
	_ = sort.Ints
	_ = strconv.ParseInt
	_ = strings.Builder{}
	_ = sync.Pool{}
	_ = time.Time{}

	_ = errors.Is
	_ = jx.Null
	_ = uuid.UUID{}
	_ = otel.GetTracerProvider
	_ = attribute.KeyValue{}
	_ = codes.Unset
	_ = metric.MeterConfig{}
	_ = syncint64.Counter(nil)
	_ = nonrecording.NewNoopMeterProvider
	_ = trace.TraceIDFromHex

	_ = conv.ToInt32
	_ = ht.NewRequest
	_ = json.Marshal
	_ = otelogen.Version
	_ = uri.PathEncoder{}
	_ = validate.Int{}
)

var _ Handler = UnimplementedHandler{}

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

// AddStickerToSet implements addStickerToSet operation.
//
// POST /addStickerToSet
func (UnimplementedHandler) AddStickerToSet(ctx context.Context, req AddStickerToSet) (r Result, _ error) {
	return r, ht.ErrNotImplemented
}

// AnswerCallbackQuery implements answerCallbackQuery operation.
//
// POST /answerCallbackQuery
func (UnimplementedHandler) AnswerCallbackQuery(ctx context.Context, req AnswerCallbackQuery) (r Result, _ error) {
	return r, ht.ErrNotImplemented
}

// AnswerInlineQuery implements answerInlineQuery operation.
//
// POST /answerInlineQuery
func (UnimplementedHandler) AnswerInlineQuery(ctx context.Context, req AnswerInlineQuery) (r Result, _ error) {
	return r, ht.ErrNotImplemented
}

// AnswerPreCheckoutQuery implements answerPreCheckoutQuery operation.
//
// POST /answerPreCheckoutQuery
func (UnimplementedHandler) AnswerPreCheckoutQuery(ctx context.Context, req AnswerPreCheckoutQuery) (r Result, _ error) {
	return r, ht.ErrNotImplemented
}

// AnswerShippingQuery implements answerShippingQuery operation.
//
// POST /answerShippingQuery
func (UnimplementedHandler) AnswerShippingQuery(ctx context.Context, req AnswerShippingQuery) (r Result, _ error) {
	return r, ht.ErrNotImplemented
}

// ApproveChatJoinRequest implements approveChatJoinRequest operation.
//
// POST /approveChatJoinRequest
func (UnimplementedHandler) ApproveChatJoinRequest(ctx context.Context, req ApproveChatJoinRequest) (r Result, _ error) {
	return r, ht.ErrNotImplemented
}

// BanChatMember implements banChatMember operation.
//
// POST /banChatMember
func (UnimplementedHandler) BanChatMember(ctx context.Context, req BanChatMember) (r Result, _ error) {
	return r, ht.ErrNotImplemented
}

// BanChatSenderChat implements banChatSenderChat operation.
//
// POST /banChatSenderChat
func (UnimplementedHandler) BanChatSenderChat(ctx context.Context, req BanChatSenderChat) (r Result, _ error) {
	return r, ht.ErrNotImplemented
}

// Close implements close operation.
//
// POST /close
func (UnimplementedHandler) Close(ctx context.Context) (r Result, _ error) {
	return r, ht.ErrNotImplemented
}

// CopyMessage implements copyMessage operation.
//
// POST /copyMessage
func (UnimplementedHandler) CopyMessage(ctx context.Context, req CopyMessage) (r ResultMessageId, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateChatInviteLink implements createChatInviteLink operation.
//
// POST /createChatInviteLink
func (UnimplementedHandler) CreateChatInviteLink(ctx context.Context, req CreateChatInviteLink) (r ResultChatInviteLink, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateNewStickerSet implements createNewStickerSet operation.
//
// POST /createNewStickerSet
func (UnimplementedHandler) CreateNewStickerSet(ctx context.Context, req CreateNewStickerSet) (r Result, _ error) {
	return r, ht.ErrNotImplemented
}

// DeclineChatJoinRequest implements declineChatJoinRequest operation.
//
// POST /declineChatJoinRequest
func (UnimplementedHandler) DeclineChatJoinRequest(ctx context.Context, req DeclineChatJoinRequest) (r Result, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteChatPhoto implements deleteChatPhoto operation.
//
// POST /deleteChatPhoto
func (UnimplementedHandler) DeleteChatPhoto(ctx context.Context, req DeleteChatPhoto) (r Result, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteChatStickerSet implements deleteChatStickerSet operation.
//
// POST /deleteChatStickerSet
func (UnimplementedHandler) DeleteChatStickerSet(ctx context.Context, req DeleteChatStickerSet) (r Result, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteMessage implements deleteMessage operation.
//
// POST /deleteMessage
func (UnimplementedHandler) DeleteMessage(ctx context.Context, req DeleteMessage) (r Result, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteMyCommands implements deleteMyCommands operation.
//
// POST /deleteMyCommands
func (UnimplementedHandler) DeleteMyCommands(ctx context.Context, req OptDeleteMyCommands) (r Result, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteStickerFromSet implements deleteStickerFromSet operation.
//
// POST /deleteStickerFromSet
func (UnimplementedHandler) DeleteStickerFromSet(ctx context.Context, req DeleteStickerFromSet) (r Result, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteWebhook implements deleteWebhook operation.
//
// POST /deleteWebhook
func (UnimplementedHandler) DeleteWebhook(ctx context.Context, req OptDeleteWebhook) (r Result, _ error) {
	return r, ht.ErrNotImplemented
}

// EditChatInviteLink implements editChatInviteLink operation.
//
// POST /editChatInviteLink
func (UnimplementedHandler) EditChatInviteLink(ctx context.Context, req EditChatInviteLink) (r ResultChatInviteLink, _ error) {
	return r, ht.ErrNotImplemented
}

// EditMessageCaption implements editMessageCaption operation.
//
// POST /editMessageCaption
func (UnimplementedHandler) EditMessageCaption(ctx context.Context, req EditMessageCaption) (r ResultMessageOrBoolean, _ error) {
	return r, ht.ErrNotImplemented
}

// EditMessageLiveLocation implements editMessageLiveLocation operation.
//
// POST /editMessageLiveLocation
func (UnimplementedHandler) EditMessageLiveLocation(ctx context.Context, req EditMessageLiveLocation) (r ResultMessageOrBoolean, _ error) {
	return r, ht.ErrNotImplemented
}

// EditMessageMedia implements editMessageMedia operation.
//
// POST /editMessageMedia
func (UnimplementedHandler) EditMessageMedia(ctx context.Context, req EditMessageMedia) (r ResultMessageOrBoolean, _ error) {
	return r, ht.ErrNotImplemented
}

// EditMessageReplyMarkup implements editMessageReplyMarkup operation.
//
// POST /editMessageReplyMarkup
func (UnimplementedHandler) EditMessageReplyMarkup(ctx context.Context, req EditMessageReplyMarkup) (r ResultMessageOrBoolean, _ error) {
	return r, ht.ErrNotImplemented
}

// EditMessageText implements editMessageText operation.
//
// POST /editMessageText
func (UnimplementedHandler) EditMessageText(ctx context.Context, req EditMessageText) (r ResultMessageOrBoolean, _ error) {
	return r, ht.ErrNotImplemented
}

// ExportChatInviteLink implements exportChatInviteLink operation.
//
// POST /exportChatInviteLink
func (UnimplementedHandler) ExportChatInviteLink(ctx context.Context, req ExportChatInviteLink) (r ResultString, _ error) {
	return r, ht.ErrNotImplemented
}

// ForwardMessage implements forwardMessage operation.
//
// POST /forwardMessage
func (UnimplementedHandler) ForwardMessage(ctx context.Context, req ForwardMessage) (r ResultMessage, _ error) {
	return r, ht.ErrNotImplemented
}

// GetChat implements getChat operation.
//
// POST /getChat
func (UnimplementedHandler) GetChat(ctx context.Context, req GetChat) (r ResultChat, _ error) {
	return r, ht.ErrNotImplemented
}

// GetChatAdministrators implements getChatAdministrators operation.
//
// POST /getChatAdministrators
func (UnimplementedHandler) GetChatAdministrators(ctx context.Context, req GetChatAdministrators) (r ResultArrayOfChatMember, _ error) {
	return r, ht.ErrNotImplemented
}

// GetChatMember implements getChatMember operation.
//
// POST /getChatMember
func (UnimplementedHandler) GetChatMember(ctx context.Context, req GetChatMember) (r ResultChatMember, _ error) {
	return r, ht.ErrNotImplemented
}

// GetChatMemberCount implements getChatMemberCount operation.
//
// POST /getChatMemberCount
func (UnimplementedHandler) GetChatMemberCount(ctx context.Context, req GetChatMemberCount) (r ResultInt, _ error) {
	return r, ht.ErrNotImplemented
}

// GetFile implements getFile operation.
//
// POST /getFile
func (UnimplementedHandler) GetFile(ctx context.Context, req GetFile) (r ResultFile, _ error) {
	return r, ht.ErrNotImplemented
}

// GetGameHighScores implements getGameHighScores operation.
//
// POST /getGameHighScores
func (UnimplementedHandler) GetGameHighScores(ctx context.Context, req GetGameHighScores) (r ResultArrayOfGameHighScore, _ error) {
	return r, ht.ErrNotImplemented
}

// GetMe implements getMe operation.
//
// POST /getMe
func (UnimplementedHandler) GetMe(ctx context.Context) (r ResultUser, _ error) {
	return r, ht.ErrNotImplemented
}

// GetMyCommands implements getMyCommands operation.
//
// POST /getMyCommands
func (UnimplementedHandler) GetMyCommands(ctx context.Context, req OptGetMyCommands) (r ResultArrayOfBotCommand, _ error) {
	return r, ht.ErrNotImplemented
}

// GetStickerSet implements getStickerSet operation.
//
// POST /getStickerSet
func (UnimplementedHandler) GetStickerSet(ctx context.Context, req GetStickerSet) (r ResultStickerSet, _ error) {
	return r, ht.ErrNotImplemented
}

// GetUpdates implements getUpdates operation.
//
// POST /getUpdates
func (UnimplementedHandler) GetUpdates(ctx context.Context, req OptGetUpdates) (r ResultArrayOfUpdate, _ error) {
	return r, ht.ErrNotImplemented
}

// GetUserProfilePhotos implements getUserProfilePhotos operation.
//
// POST /getUserProfilePhotos
func (UnimplementedHandler) GetUserProfilePhotos(ctx context.Context, req GetUserProfilePhotos) (r ResultUserProfilePhotos, _ error) {
	return r, ht.ErrNotImplemented
}

// GetWebhookInfo implements getWebhookInfo operation.
//
// POST /getWebhookInfo
func (UnimplementedHandler) GetWebhookInfo(ctx context.Context) (r ResultWebhookInfo, _ error) {
	return r, ht.ErrNotImplemented
}

// LeaveChat implements leaveChat operation.
//
// POST /leaveChat
func (UnimplementedHandler) LeaveChat(ctx context.Context, req LeaveChat) (r Result, _ error) {
	return r, ht.ErrNotImplemented
}

// LogOut implements logOut operation.
//
// POST /logOut
func (UnimplementedHandler) LogOut(ctx context.Context) (r Result, _ error) {
	return r, ht.ErrNotImplemented
}

// PinChatMessage implements pinChatMessage operation.
//
// POST /pinChatMessage
func (UnimplementedHandler) PinChatMessage(ctx context.Context, req PinChatMessage) (r Result, _ error) {
	return r, ht.ErrNotImplemented
}

// PromoteChatMember implements promoteChatMember operation.
//
// POST /promoteChatMember
func (UnimplementedHandler) PromoteChatMember(ctx context.Context, req PromoteChatMember) (r Result, _ error) {
	return r, ht.ErrNotImplemented
}

// RestrictChatMember implements restrictChatMember operation.
//
// POST /restrictChatMember
func (UnimplementedHandler) RestrictChatMember(ctx context.Context, req RestrictChatMember) (r Result, _ error) {
	return r, ht.ErrNotImplemented
}

// RevokeChatInviteLink implements revokeChatInviteLink operation.
//
// POST /revokeChatInviteLink
func (UnimplementedHandler) RevokeChatInviteLink(ctx context.Context, req RevokeChatInviteLink) (r ResultChatInviteLink, _ error) {
	return r, ht.ErrNotImplemented
}

// SendAnimation implements sendAnimation operation.
//
// POST /sendAnimation
func (UnimplementedHandler) SendAnimation(ctx context.Context, req SendAnimation) (r ResultMessage, _ error) {
	return r, ht.ErrNotImplemented
}

// SendAudio implements sendAudio operation.
//
// POST /sendAudio
func (UnimplementedHandler) SendAudio(ctx context.Context, req SendAudio) (r ResultMessage, _ error) {
	return r, ht.ErrNotImplemented
}

// SendChatAction implements sendChatAction operation.
//
// POST /sendChatAction
func (UnimplementedHandler) SendChatAction(ctx context.Context, req SendChatAction) (r Result, _ error) {
	return r, ht.ErrNotImplemented
}

// SendContact implements sendContact operation.
//
// POST /sendContact
func (UnimplementedHandler) SendContact(ctx context.Context, req SendContact) (r ResultMessage, _ error) {
	return r, ht.ErrNotImplemented
}

// SendDice implements sendDice operation.
//
// POST /sendDice
func (UnimplementedHandler) SendDice(ctx context.Context, req SendDice) (r ResultMessage, _ error) {
	return r, ht.ErrNotImplemented
}

// SendDocument implements sendDocument operation.
//
// POST /sendDocument
func (UnimplementedHandler) SendDocument(ctx context.Context, req SendDocument) (r ResultMessage, _ error) {
	return r, ht.ErrNotImplemented
}

// SendGame implements sendGame operation.
//
// POST /sendGame
func (UnimplementedHandler) SendGame(ctx context.Context, req SendGame) (r ResultMessage, _ error) {
	return r, ht.ErrNotImplemented
}

// SendInvoice implements sendInvoice operation.
//
// POST /sendInvoice
func (UnimplementedHandler) SendInvoice(ctx context.Context, req SendInvoice) (r ResultMessage, _ error) {
	return r, ht.ErrNotImplemented
}

// SendLocation implements sendLocation operation.
//
// POST /sendLocation
func (UnimplementedHandler) SendLocation(ctx context.Context, req SendLocation) (r ResultMessage, _ error) {
	return r, ht.ErrNotImplemented
}

// SendMediaGroup implements sendMediaGroup operation.
//
// POST /sendMediaGroup
func (UnimplementedHandler) SendMediaGroup(ctx context.Context, req SendMediaGroup) (r ResultArrayOfMessage, _ error) {
	return r, ht.ErrNotImplemented
}

// SendMessage implements sendMessage operation.
//
// POST /sendMessage
func (UnimplementedHandler) SendMessage(ctx context.Context, req SendMessage) (r ResultMessage, _ error) {
	return r, ht.ErrNotImplemented
}

// SendPhoto implements sendPhoto operation.
//
// POST /sendPhoto
func (UnimplementedHandler) SendPhoto(ctx context.Context, req SendPhoto) (r ResultMessage, _ error) {
	return r, ht.ErrNotImplemented
}

// SendPoll implements sendPoll operation.
//
// POST /sendPoll
func (UnimplementedHandler) SendPoll(ctx context.Context, req SendPoll) (r ResultMessage, _ error) {
	return r, ht.ErrNotImplemented
}

// SendSticker implements sendSticker operation.
//
// POST /sendSticker
func (UnimplementedHandler) SendSticker(ctx context.Context, req SendSticker) (r ResultMessage, _ error) {
	return r, ht.ErrNotImplemented
}

// SendVenue implements sendVenue operation.
//
// POST /sendVenue
func (UnimplementedHandler) SendVenue(ctx context.Context, req SendVenue) (r ResultMessage, _ error) {
	return r, ht.ErrNotImplemented
}

// SendVideo implements sendVideo operation.
//
// POST /sendVideo
func (UnimplementedHandler) SendVideo(ctx context.Context, req SendVideo) (r ResultMessage, _ error) {
	return r, ht.ErrNotImplemented
}

// SendVideoNote implements sendVideoNote operation.
//
// POST /sendVideoNote
func (UnimplementedHandler) SendVideoNote(ctx context.Context, req SendVideoNote) (r ResultMessage, _ error) {
	return r, ht.ErrNotImplemented
}

// SendVoice implements sendVoice operation.
//
// POST /sendVoice
func (UnimplementedHandler) SendVoice(ctx context.Context, req SendVoice) (r ResultMessage, _ error) {
	return r, ht.ErrNotImplemented
}

// SetChatAdministratorCustomTitle implements setChatAdministratorCustomTitle operation.
//
// POST /setChatAdministratorCustomTitle
func (UnimplementedHandler) SetChatAdministratorCustomTitle(ctx context.Context, req SetChatAdministratorCustomTitle) (r Result, _ error) {
	return r, ht.ErrNotImplemented
}

// SetChatDescription implements setChatDescription operation.
//
// POST /setChatDescription
func (UnimplementedHandler) SetChatDescription(ctx context.Context, req SetChatDescription) (r Result, _ error) {
	return r, ht.ErrNotImplemented
}

// SetChatPermissions implements setChatPermissions operation.
//
// POST /setChatPermissions
func (UnimplementedHandler) SetChatPermissions(ctx context.Context, req SetChatPermissions) (r Result, _ error) {
	return r, ht.ErrNotImplemented
}

// SetChatPhoto implements setChatPhoto operation.
//
// POST /setChatPhoto
func (UnimplementedHandler) SetChatPhoto(ctx context.Context, req SetChatPhoto) (r Result, _ error) {
	return r, ht.ErrNotImplemented
}

// SetChatStickerSet implements setChatStickerSet operation.
//
// POST /setChatStickerSet
func (UnimplementedHandler) SetChatStickerSet(ctx context.Context, req SetChatStickerSet) (r Result, _ error) {
	return r, ht.ErrNotImplemented
}

// SetChatTitle implements setChatTitle operation.
//
// POST /setChatTitle
func (UnimplementedHandler) SetChatTitle(ctx context.Context, req SetChatTitle) (r Result, _ error) {
	return r, ht.ErrNotImplemented
}

// SetGameScore implements setGameScore operation.
//
// POST /setGameScore
func (UnimplementedHandler) SetGameScore(ctx context.Context, req SetGameScore) (r Result, _ error) {
	return r, ht.ErrNotImplemented
}

// SetMyCommands implements setMyCommands operation.
//
// POST /setMyCommands
func (UnimplementedHandler) SetMyCommands(ctx context.Context, req SetMyCommands) (r Result, _ error) {
	return r, ht.ErrNotImplemented
}

// SetPassportDataErrors implements setPassportDataErrors operation.
//
// POST /setPassportDataErrors
func (UnimplementedHandler) SetPassportDataErrors(ctx context.Context, req SetPassportDataErrors) (r Result, _ error) {
	return r, ht.ErrNotImplemented
}

// SetStickerPositionInSet implements setStickerPositionInSet operation.
//
// POST /setStickerPositionInSet
func (UnimplementedHandler) SetStickerPositionInSet(ctx context.Context, req SetStickerPositionInSet) (r Result, _ error) {
	return r, ht.ErrNotImplemented
}

// SetStickerSetThumb implements setStickerSetThumb operation.
//
// POST /setStickerSetThumb
func (UnimplementedHandler) SetStickerSetThumb(ctx context.Context, req SetStickerSetThumb) (r Result, _ error) {
	return r, ht.ErrNotImplemented
}

// SetWebhook implements setWebhook operation.
//
// POST /setWebhook
func (UnimplementedHandler) SetWebhook(ctx context.Context, req SetWebhook) (r Result, _ error) {
	return r, ht.ErrNotImplemented
}

// StopMessageLiveLocation implements stopMessageLiveLocation operation.
//
// POST /stopMessageLiveLocation
func (UnimplementedHandler) StopMessageLiveLocation(ctx context.Context, req StopMessageLiveLocation) (r ResultMessageOrBoolean, _ error) {
	return r, ht.ErrNotImplemented
}

// StopPoll implements stopPoll operation.
//
// POST /stopPoll
func (UnimplementedHandler) StopPoll(ctx context.Context, req StopPoll) (r ResultPoll, _ error) {
	return r, ht.ErrNotImplemented
}

// UnbanChatMember implements unbanChatMember operation.
//
// POST /unbanChatMember
func (UnimplementedHandler) UnbanChatMember(ctx context.Context, req UnbanChatMember) (r Result, _ error) {
	return r, ht.ErrNotImplemented
}

// UnbanChatSenderChat implements unbanChatSenderChat operation.
//
// POST /unbanChatSenderChat
func (UnimplementedHandler) UnbanChatSenderChat(ctx context.Context, req UnbanChatSenderChat) (r Result, _ error) {
	return r, ht.ErrNotImplemented
}

// UnpinAllChatMessages implements unpinAllChatMessages operation.
//
// POST /unpinAllChatMessages
func (UnimplementedHandler) UnpinAllChatMessages(ctx context.Context, req UnpinAllChatMessages) (r Result, _ error) {
	return r, ht.ErrNotImplemented
}

// UnpinChatMessage implements unpinChatMessage operation.
//
// POST /unpinChatMessage
func (UnimplementedHandler) UnpinChatMessage(ctx context.Context, req UnpinChatMessage) (r Result, _ error) {
	return r, ht.ErrNotImplemented
}

// UploadStickerFile implements uploadStickerFile operation.
//
// POST /uploadStickerFile
func (UnimplementedHandler) UploadStickerFile(ctx context.Context, req UploadStickerFile) (r ResultFile, _ error) {
	return r, ht.ErrNotImplemented
}

// NewError creates ErrorStatusCode from error returned by handler.
//
// Used for common default response.
func (UnimplementedHandler) NewError(ctx context.Context, err error) (r ErrorStatusCode) {
	return r
}
