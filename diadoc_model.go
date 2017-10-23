package crawler

import (

	//"github.com/satori/go.uuid"
	"time"
)

const ERR_INVOICE_SIGN_NOT_FOUND = "can't find signature for invoice"

type (
	SignedContent struct {
		Content   string `json:"Content"`
		Signature string `json:"Signature"`
	}
	SignerCorrection struct {
		ErrorMessage string `json:"ErrorMessage"`
		Signer       Signer
	}
	XmlSignatureRejections struct {
		BoxID                  string `json:"BoxId"`
		MessageID              string `json:"MessageId"`
		XmlSignatureRejections []Buyers
	}
	Buyers struct {
		ParentEntityID string `json:"ParentEntityId"`
		SignedContent  SignedContent
	}
	Receipts struct {
		ParentEntityID string `json:"ParentEntityId"`
		SignedContent  SignedContent
		Comment        string `json:"Comment"`
	}
	BuyerTitle struct {
		DocumentCreator  string `json:"DocumentCreator"`
		OperationContent string `json:"OperationContent"`
		Signers          []Signers
	}

	Signers struct {
		BoxID         string `json:"BoxId"`
		SignerDetails SignerDetails
	}
	UniversalTransferDocumentBuyer struct {
		BoxID                                string `json:"BoxId"`
		MessageID                            string `json:"MessageId"`
		UniversalTransferDocumentBuyerTitles []Buyers
	}
	ReceiptAttachment struct {
		BoxID     string `json:"BoxId"`
		MessageID string `json:"MessageId"`
		Receipts  []Receipts
	}

	Signer struct {
		SignerCertificate           []byte `json:"SignerCertificate"`
		SignerDetails               SignerDetailsForCorrect
		SignerCertificateThumbprint string `json:"SignerCertificateThumbprint"`
	}

	SignerDetailsForCorrect struct {
		Surname                               string `json:"Surname"`
		FirstName                             string `json:"FirstName"`
		Patronymic                            string `json:"Patronymic"`
		JobTitle                              string `json:"JobTitle"`
		Inn                                   string `json:"Inn"`
		SoleProprietorRegistrationCertificate string `json:"SoleProprietorRegistrationCertificate"`
	}

	Documents struct {
		IndexKey               string        `json:"IndexKey"`
		MessageID              string        `json:"MessageId"`
		EntityID               string        `json:"EntityId"`
		CreationTimestampTicks int64         `json:"CreationTimestampTicks"`
		CounteragentBoxID      string        `json:"CounteragentBoxId"`
		DocumentType           string        `json:"DocumentType"`
		InitialDocumentIds     []interface{} `json:"InitialDocumentIds"`
		SubordinateDocumentIds []interface{} `json:"SubordinateDocumentIds"`
		Content                struct {
			Size int `json:"Size"`
		} `json:"Content"`
		FileName                          string        `json:"FileName"`
		DocumentDate                      string        `json:"DocumentDate"`
		DocumentNumber                    string        `json:"DocumentNumber"`
		IsDeleted                         bool          `json:"IsDeleted"`
		DepartmentID                      string        `json:"DepartmentId"`
		IsTest                            bool          `json:"IsTest"`
		FromDepartmentID                  string        `json:"FromDepartmentId"`
		ToDepartmentID                    string        `json:"ToDepartmentId"`
		RevocationStatus                  string        `json:"RevocationStatus"`
		SendTimestampTicks                int64         `json:"SendTimestampTicks"`
		DeliveryTimestampTicks            int64         `json:"DeliveryTimestampTicks"`
		ForwardDocumentEvents             []interface{} `json:"ForwardDocumentEvents"`
		RoamingNotificationStatus         string        `json:"RoamingNotificationStatus"`
		HasCustomPrintForm                bool          `json:"HasCustomPrintForm"`
		CustomData                        []interface{} `json:"CustomData"`
		DocumentDirection                 string        `json:"DocumentDirection"`
		LastModificationTimestampTicks    int64         `json:"LastModificationTimestampTicks"`
		IsEncryptedContent                bool          `json:"IsEncryptedContent"`
		SenderSignatureStatus             string        `json:"SenderSignatureStatus"`
		IsRead                            bool          `json:"IsRead"`
		PacketIsLocked                    bool          `json:"PacketIsLocked"`
		UniversalTransferDocumentMetadata struct {
			DocumentStatus            string `json:"DocumentStatus"`
			Total                     string `json:"Total"`
			Vat                       string `json:"Vat"`
			Grounds                   string `json:"Grounds"`
			DocumentFunction          string `json:"DocumentFunction"`
			Currency                  int    `json:"Currency"`
			ConfirmationDateTimeTicks int    `json:"ConfirmationDateTimeTicks"`
			InvoiceAmendmentFlags     int    `json:"InvoiceAmendmentFlags"`
		} `json:"UniversalTransferDocumentMetadata"`
		ResolutionRouteID     string    `json:"ResolutionRouteId"`
		AttachmentVersion     string    `json:"AttachmentVersion"`
		ProxySignatureStatus  string    `json:"ProxySignatureStatus"`
		MessageIDGUID         string    `json:"MessageIdGuid"`
		EntityIDGUID          string    `json:"EntityIdGuid"`
		CreationTimestamp     time.Time `json:"CreationTimestamp"`
		CounteragentBoxIDGUID string    `json:"CounteragentBoxIdGuid"`
	}
	Docs struct {
		TotalCount int `json:"TotalCount"`
		Documents  []Documents
	}

	DocumentMessage struct {
		MessageID               string `json:"MessageId"`
		TimestampTicks          int64  `json:"TimestampTicks"`
		LastPatchTimestampTicks int64  `json:"LastPatchTimestampTicks"`
		FromBoxID               string `json:"FromBoxId"`
		FromTitle               string `json:"FromTitle"`
		ToBoxID                 string `json:"ToBoxId"`
		ToTitle                 string `json:"ToTitle"`
		Entities                []struct {
			EntityType     string `json:"EntityType"`
			EntityID       string `json:"EntityId"`
			ParentEntityID string `json:"ParentEntityId"`
			Content        struct {
				Size int    `json:"Size"`
				Data string `json:"Data"`
			} `json:"Content"`
			AttachmentType         string `json:"AttachmentType"`
			FileName               string `json:"FileName"`
			NeedRecipientSignature bool   `json:"NeedRecipientSignature"`
			SignerBoxID            string `json:"SignerBoxId"`
			NotDeliveredEventID    string `json:"NotDeliveredEventId"`
			DocumentInfo           struct {
				MessageID              string        `json:"MessageId"`
				EntityID               string        `json:"EntityId"`
				CreationTimestampTicks int64         `json:"CreationTimestampTicks"`
				CounteragentBoxID      string        `json:"CounteragentBoxId"`
				DocumentType           string        `json:"DocumentType"`
				InitialDocumentIds     []interface{} `json:"InitialDocumentIds"`
				SubordinateDocumentIds []interface{} `json:"SubordinateDocumentIds"`
				Content                struct {
					Size int `json:"Size"`
				} `json:"Content"`
				FileName                          string        `json:"FileName"`
				DocumentDate                      string        `json:"DocumentDate"`
				DocumentNumber                    string        `json:"DocumentNumber"`
				IsDeleted                         bool          `json:"IsDeleted"`
				DepartmentID                      string        `json:"DepartmentId"`
				IsTest                            bool          `json:"IsTest"`
				FromDepartmentID                  string        `json:"FromDepartmentId"`
				ToDepartmentID                    string        `json:"ToDepartmentId"`
				CustomDocumentID                  string        `json:"CustomDocumentId"`
				RevocationStatus                  string        `json:"RevocationStatus"`
				SendTimestampTicks                int64         `json:"SendTimestampTicks"`
				DeliveryTimestampTicks            int64         `json:"DeliveryTimestampTicks"`
				ForwardDocumentEvents             []interface{} `json:"ForwardDocumentEvents"`
				RoamingNotificationStatus         string        `json:"RoamingNotificationStatus"`
				HasCustomPrintForm                bool          `json:"HasCustomPrintForm"`
				CustomData                        []interface{} `json:"CustomData"`
				DocumentDirection                 string        `json:"DocumentDirection"`
				LastModificationTimestampTicks    int64         `json:"LastModificationTimestampTicks"`
				IsEncryptedContent                bool          `json:"IsEncryptedContent"`
				SenderSignatureStatus             string        `json:"SenderSignatureStatus"`
				IsRead                            bool          `json:"IsRead"`
				PacketIsLocked                    bool          `json:"PacketIsLocked"`
				UniversalTransferDocumentMetadata struct {
					DocumentStatus            string `json:"DocumentStatus"`
					Total                     string `json:"Total"`
					Vat                       string `json:"Vat"`
					Grounds                   string `json:"Grounds"`
					DocumentFunction          string `json:"DocumentFunction"`
					Currency                  int    `json:"Currency"`
					ConfirmationDateTimeTicks int    `json:"ConfirmationDateTimeTicks"`
					InvoiceAmendmentFlags     int    `json:"InvoiceAmendmentFlags"`
				} `json:"UniversalTransferDocumentMetadata"`
				ResolutionRouteID     string    `json:"ResolutionRouteId"`
				AttachmentVersion     string    `json:"AttachmentVersion"`
				ProxySignatureStatus  string    `json:"ProxySignatureStatus"`
				MessageIDGUID         string    `json:"MessageIdGuid"`
				EntityIDGUID          string    `json:"EntityIdGuid"`
				CreationTimestamp     time.Time `json:"CreationTimestamp"`
				CounteragentBoxIDGUID string    `json:"CounteragentBoxIdGuid"`
			} `json:"DocumentInfo,omitempty"`
			RawCreationDate        int64  `json:"RawCreationDate"`
			SignerDepartmentID     string `json:"SignerDepartmentId"`
			NeedReceipt            bool   `json:"NeedReceipt"`
			IsApprovementSignature bool   `json:"IsApprovementSignature"`
			IsEncryptedContent     bool   `json:"IsEncryptedContent"`
			AttachmentVersion      string `json:"AttachmentVersion"`
			PacketID               string `json:"PacketId,omitempty"`
		} `json:"Entities"`
		IsDraft                           bool          `json:"IsDraft"`
		DraftIsLocked                     bool          `json:"DraftIsLocked"`
		DraftIsRecycled                   bool          `json:"DraftIsRecycled"`
		CreatedFromDraftID                string        `json:"CreatedFromDraftId"`
		DraftIsTransformedToMessageIDList []interface{} `json:"DraftIsTransformedToMessageIdList"`
		IsTest                            bool          `json:"IsTest"`
		IsInternal                        bool          `json:"IsInternal"`
		IsProxified                       bool          `json:"IsProxified"`
		ProxyBoxID                        string        `json:"ProxyBoxId"`
		ProxyTitle                        string        `json:"ProxyTitle"`
		PacketIsLocked                    bool          `json:"PacketIsLocked"`
	}

	SignerDetails struct {
		Surname                               string `json:"Surname"`
		FirstName                             string `json:"FirstName"`
		Patronymic                            string `json:"Patronymic"`
		JobTitle                              string `json:"JobTitle"`
		Inn                                   string `json:"Inn"`
		SignerType                            int    `json:"SignerType"`
		SignerOrganizationName                string `json:"SignerOrganizationName"`
		SignerInfo                            string `json:"SignerInfo"`
		SignerPowers                          int    `json:"SignerPowers"`
		SignerStatus                          int    `json:"SignerStatus"`
		SignerPowersBase                      string `json:"SignerPowersBase"`
		SignerOrgPowersBase                   string `json:"SignerOrgPowersBase"`
		SoleProprietorRegistrationCertificate string `json:"SoleProprietorRegistrationCertificate"`
	}
)
