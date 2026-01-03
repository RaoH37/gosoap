package zimbraMail

type FolderView string

const (
	FolderViewAll         FolderView = ""
	FolderViewUnknown     FolderView = "unknown"
	FolderViewMessage     FolderView = "message"
	FolderViewAppointment FolderView = "appointment"
	FolderViewContact     FolderView = "contact"
	FolderViewDocument    FolderView = "document"
	FolderViewTask        FolderView = "task"
)

type ItemOperation string

const (
	ItemOperationUnCheck             ItemOperation = "[!]check"
	ItemOperationUnDisableactivesync ItemOperation = "[!]disableactivesync"
	ItemOperationUnGrant             ItemOperation = "[!]grant"
	ItemOperationUnSyncon            ItemOperation = "[!]syncon"
	ItemOperationAcceptsubsreq       ItemOperation = "acceptSubsReq"
	ItemOperationAddmembers          ItemOperation = "addMembers"
	ItemOperationAddowners           ItemOperation = "addOwners"
	ItemOperationColor               ItemOperation = "color"
	ItemOperationCopy                ItemOperation = "copy"
	ItemOperationDelete              ItemOperation = "delete"
	ItemOperationDumpsterdelete      ItemOperation = "dumpsterdelete"
	ItemOperationEmpty               ItemOperation = "empty"
	ItemOperationFb                  ItemOperation = "fb"
	ItemOperationFlag                ItemOperation = "flag"
	ItemOperationGrantrights         ItemOperation = "grantRights"
	ItemOperationImport              ItemOperation = "import"
	ItemOperationLock                ItemOperation = "lock"
	ItemOperationModify              ItemOperation = "modify"
	ItemOperationMove                ItemOperation = "move"
	ItemOperationMute                ItemOperation = "mute"
	ItemOperationPriority            ItemOperation = "priority"
	ItemOperationRead                ItemOperation = "read"
	ItemOperationRecover             ItemOperation = "recover"
	ItemOperationRejectsubsreq       ItemOperation = "rejectSubsReq"
	ItemOperationRemovemembers       ItemOperation = "removeMembers"
	ItemOperationRemoveowners        ItemOperation = "removeOwners"
	ItemOperationRename              ItemOperation = "rename"
	ItemOperationResetimapuid        ItemOperation = "resetimapuid"
	ItemOperationRetentionpolicy     ItemOperation = "retentionpolicy"
	ItemOperationRevokerights        ItemOperation = "revokeRights"
	ItemOperationRevokeorphangrants  ItemOperation = "revokeorphangrants"
	ItemOperationSetowners           ItemOperation = "setOwners"
	ItemOperationSetrights           ItemOperation = "setRights"
	ItemOperationSpam                ItemOperation = "spam"
	ItemOperationSync                ItemOperation = "sync"
	ItemOperationTag                 ItemOperation = "tag"
	ItemOperationTrash               ItemOperation = "trash"
	ItemOperationUnlock              ItemOperation = "unlock"
	ItemOperationUpdate              ItemOperation = "update"
	ItemOperationUrl                 ItemOperation = "url"
	ItemOperationWebofflinesyncdays  ItemOperation = "webofflinesyncdays"
)
