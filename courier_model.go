package crawler

const (
	UPD_SF_DOP_NAME string = "UPD_SF_DOP"
	UPD_SF_NAME     string = "UPD_SF"
	UPD_DOP_NAME    string = "UPD_DOP"

	ReceiveConfirmationNotificationStringPart string = "DP_PDOTPR"

	//document processing
	DocumentTitle int = 1
	SendConfirmation int = 3
	ReceiveNotificationOfSendConfirmationOfReceiveNotificationOfDocument int = 8
	ReceiveNotificationOfSendConfirmation int = 7
	ReceiveNotificationOfDocument int = 5
	SendConfirmationOfReceiveNotificationOfDocument int = 6

	//document closing
	ClarificationNotice int = 9
	ClarificationNoticeName string = "XmlSignatureRejection"
	AcceptDocument int = 11
	AcceptDocumentName string = "UniversalTransferDocumentBuyerTitle"
)

type (
	AuthCreds struct {
		Login    string `json:"username"`
		Password string `json:"password"`
	}

	AuthResp struct {
		Token string `json:"token"`
	}
	DocumentAvoidanceComment struct {
		Comment string `json:"comment"`
	}

	CourierTickets []struct {
		Content struct {
			Filename string `json:"filename"`
			MimeType string `json:"mimeType"`
			Content  string `json:"content"`
		} `json:"content"`
		Signature []struct {
			UUID       string `json:"uuid"`
			SignerCode string `json:"signerCode"`
			Content    string `json:"content"`
		} `json:"signature"`
		NeedReply    bool        `json:"needReply"`
		ID           int         `json:"id"`
		DocumentID   int         `json:"documentId"`
		UUID         string      `json:"uuid"`
		ParentUUID   interface{} `json:"parentUuid"`
		Direction    int         `json:"direction"`
		Type         int         `json:"type"`
		ExtendedType int         `json:"extendedType"`
		Date         string      `json:"date"`
	}

	Doc struct {
		ID                  int         `json:"id"`
		TypeID              int         `json:"typeId"`
		TypeName            string      `json:"typeName"`
		DocumentTypeCode    string      `json:"documentTypeCode"`
		Description         string      `json:"description"`
		BarcodeType         int         `json:"barcodeType"`
		BarcodeValue        interface{} `json:"barcodeValue"`
		SenderID            int         `json:"senderId"`
		ReceiverID          int         `json:"receiverId"`
		ParticipantID       int         `json:"participantId"`
		ParticipantName     string      `json:"participantName"`
		Status              int         `json:"status"`
		Date                string      `json:"date"`
		RequestSign         bool        `json:"requestSign"`
		Number              string      `json:"number"`
		NetSum              float64     `json:"netSum"`
		VatSum              float64     `json:"vatSum"`
		TotalSum            float64     `json:"totalSum"`
		Actions             int         `json:"actions"`
		IsMarked            bool        `json:"isMarked"`
		Created             string      `json:"created"`
		SendDate            string      `json:"sendDate"`
		ReceiveDate         string      `json:"receiveDate"`
		IsDocflowCompleted  bool        `json:"isDocflowCompleted"`
		Comment             interface{} `json:"comment"`
		ContractNumber      interface{} `json:"contractNumber"`
		ContractDate        interface{} `json:"contractDate"`
		Type                interface{} `json:"type"`
		FormType            int         `json:"formType"`
		SellerCode          interface{} `json:"sellerCode"`
		BuyerCode           interface{} `json:"buyerCode"`
		FormatVersion       interface{} `json:"formatVersion"`
		Filename            string      `json:"filename"`
		RevokeReason        interface{} `json:"revokeReason"`
		ContractDescription string      `json:"contractDescription"`
		IsRoaming           bool        `json:"isRoaming"`
		Content             struct {
			Filename string `json:"filename"`
			MimeType string `json:"mimeType"`
			Content  string `json:"content"`
		} `json:"content"`
		IsPrintable          bool          `json:"isPrintable"`
		Relations            []interface{} `json:"relations"`
		StatusName           string        `json:"statusName"`
		RejectReason         interface{}   `json:"rejectReason"`
		SenderName           interface{}   `json:"senderName"`
		ReceiverName         interface{}   `json:"receiverName"`
		ExistRouteSigner     bool          `json:"existRouteSigner"`
		DocumentRelationType int           `json:"documentRelationType"`
		Attachments          []interface{} `json:"attachments"`
		StatusChanged        string        `json:"statusChanged"`
		PackageID            interface{}   `json:"packageId"`
	}

	CourierDocuments []Doc

	DocumentsRequest struct {
		Folder string `json:"folder"`
	}

	DocumentTickets []struct {
		Content struct {
			Filename string `json:"filename"`
			MimeType string `json:"mimeType"`
			Content  string `json:"content"`
		} `json:"content"`
		Signature []struct {
			UUID       string `json:"uuid"`
			SignerCode string `json:"signerCode"`
			Content    string `json:"content"`
		} `json:"signature"`
		NeedReply    bool        `json:"needReply"`
		ID           int         `json:"id"`
		DocumentID   int         `json:"documentId"`
		UUID         string      `json:"uuid"`
		ParentUUID   interface{} `json:"parentUuid"`
		Direction    int         `json:"direction"`
		Type         int         `json:"type"`
		ExtendedType int         `json:"extendedType"`
		Date         string      `json:"date"`
	}
	ErrorMessage struct {
		ErrorNumber string `json:"errorNumber"`
		Message     string `json:"message"`
	}

	TicketReply struct {
		Filename string `json:"filename"`
		MimeType string `json:"mimeType"`
		Content  string `json:"content"`
	}
	TicketAddReply struct {
		Content struct {
			Filename string `json:"filename"`
			MimeType string `json:"mimeType"`
			Content  string `json:"content"`
		} `json:"content"`
		Signature []struct {
			UUID       string `json:"uuid"`
			SignerCode string `json:"signerCode"`
			Content    string `json:"content"`
		} `json:"signature"`
		NeedReply    bool        `json:"needReply"`
		ID           int         `json:"id"`
		DocumentID   interface{} `json:"documentId"`
		UUID         string      `json:"uuid"`
		ParentUUID   string      `json:"parentUuid"`
		Direction    int         `json:"direction"`
		Type         int         `json:"type"`
		ExtendedType int         `json:"extendedType"`
		Date         string      `json:"date"`
	}

	TicketAdd struct {
		Filename  string `json:"filename"`
		MimeType  string `json:"mimeType"`
		Content   string `json:"content"`
		Signature string `json:"signature"`
	}

	DocumentSigner struct {
		Signer          SignerCourier
		SignerAuthority SignerAuthority
	}
	SignerAuthority struct {
		AuthorityScope int    `json:"authorityScope"`
		Condition      int    `json:"condition"`
		Authority      string `json:"authority"`
		OrgAuthority   string `json:"orgAuthority"`
	}
	SignerCourier struct {
		SignerType             int    `json:"signerType"`
		Inn                    string `json:"inn"`
		Name                   string `json:"name"`
		RegistrationRequisites string `json:"registrationRequisites"`
		Person                 Person
		Title                  string `json:"title"`
	}
	Person struct {
		LastName   string `json:"lastName"`
		FirstName  string `json:"firstName"`
		MiddleName string `json:"middleName"`
	}
	AcceptTitle struct {
		Date           string `json:"date"`
		DealSummary    string `json:"dealSummary"`
		DocumentSigner []DocumentSigner
	}

	DocumentDetails struct {
		ID                  int         `json:"id"`
		TypeID              int         `json:"typeId"`
		TypeName            string      `json:"typeName"`
		DocumentTypeCode    string      `json:"documentTypeCode"`
		Description         string      `json:"description"`
		BarcodeType         int         `json:"barcodeType"`
		BarcodeValue        interface{} `json:"barcodeValue"`
		SenderID            int         `json:"senderId"`
		ReceiverID          int         `json:"receiverId"`
		ParticipantID       int         `json:"participantId"`
		ParticipantName     interface{} `json:"participantName"`
		Status              int         `json:"status"`
		Date                string      `json:"date"`
		RequestSign         bool        `json:"requestSign"`
		Number              string      `json:"number"`
		NetSum              float64     `json:"netSum"`
		VatSum              float64     `json:"vatSum"`
		TotalSum            float64     `json:"totalSum"`
		Actions             int         `json:"actions"`
		IsMarked            bool        `json:"isMarked"`
		Created             string      `json:"created"`
		SendDate            string      `json:"sendDate"`
		ReceiveDate         string      `json:"receiveDate"`
		IsDocflowCompleted  bool        `json:"isDocflowCompleted"`
		Comment             interface{} `json:"comment"`
		ContractNumber      interface{} `json:"contractNumber"`
		ContractDate        interface{} `json:"contractDate"`
		Type                int         `json:"type"`
		FormType            int         `json:"formType"`
		SellerCode          string      `json:"sellerCode"`
		BuyerCode           string      `json:"buyerCode"`
		FormatVersion       string      `json:"formatVersion"`
		Filename            string      `json:"filename"`
		RevokeReason        interface{} `json:"revokeReason"`
		ContractDescription string      `json:"contractDescription"`
		IsRoaming           bool        `json:"isRoaming"`
		Content             struct {
			Filename string `json:"filename"`
			MimeType string `json:"mimeType"`
			Content  string `json:"content"`
		} `json:"content"`
		IsPrintable bool `json:"isPrintable"`
		Relations   []struct {
			ID          int         `json:"id"`
			ParentID    interface{} `json:"parentId"`
			Relation    interface{} `json:"relation"`
			Description string      `json:"description"`
		} `json:"relations"`
		StatusName           string        `json:"statusName"`
		RejectReason         interface{}   `json:"rejectReason"`
		SenderName           string        `json:"senderName"`
		ReceiverName         string        `json:"receiverName"`
		ExistRouteSigner     bool          `json:"existRouteSigner"`
		DocumentRelationType int           `json:"documentRelationType"`
		Attachments          []interface{} `json:"attachments"`
		StatusChanged        string        `json:"statusChanged"`
		PackageID            interface{}   `json:"packageId"`
	}
)
