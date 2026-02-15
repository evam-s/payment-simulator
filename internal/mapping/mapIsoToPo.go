package mapping

import (
	// "encoding/json"
	// "errors"
	"log"
	"payment-simulator/internal/iso20022/isomodels"
	"payment-simulator/internal/models"

	"github.com/gin-gonic/gin"
)

func MapXmlPacs008(c *gin.Context) (*isomodels.Pacs008, error) {
	var isoPacs isomodels.Pacs008
	if err := c.ShouldBindXML(&isoPacs); err != nil {
		log.Println("Error in Binding to XML: ", isoPacs, "Error:", err)
		return nil, err
	} else {
		return &isoPacs, nil
	}
}

func MapPacs008ToPo(isoPacs *isomodels.Pacs008) (error, *models.PaymentOrder) {

	var po models.PaymentOrder

	grpHdr := isoPacs.FIToFICstmrCdtTrf.GrpHdr

	// if len(isoPacs.FIToFICstmrCdtTrf.CdtTrfTxInf) == 0 {
	// 	return errors.New("Minimum Number of FIToFICstmrCdtTrf.CdtTrfTxInf must be 1"), nil
	// }

	txn := isoPacs.FIToFICstmrCdtTrf.CdtTrfTxInf[0]

	var charges []models.Charges

	for _, ch := range txn.ChrgsInf {
		charges = append(charges, models.Charges{
			Amount: *mapAmount(ch.Amt),
			Agent:  *mapAgent(ch.Agt),
			Type: models.ChrgsTp{
				Code:              ch.Tp.Cd,
				ProprietaryId:     ch.Tp.Prtry.Id,
				ProprietaryIssuer: ch.Tp.Prtry.Issr,
			},
		})
	}

	po = models.PaymentOrder{
		Debtor:                  *mapParty(txn.Dbtr),
		DebtorAcct:              *mapAccount(txn.DbtrAcct),
		Creditor:                *mapParty(txn.Cdtr),
		CreditorAcct:            *mapAccount(txn.CdtrAcct),
		SettlementAmount:        *mapAmount(txn.IntrBkSttlmAmt),
		InstructionId:           txn.PmtId.InstrId,
		TransactionId:           txn.PmtId.TxId,
		EndToEndId:              txn.PmtId.EndToEndId,
		UETR:                    txn.PmtId.UETR,
		ClearingSystemReference: txn.PmtId.ClrSysRef,
		MsgId:                   grpHdr.MsgId,
		CreationDateTime:        grpHdr.CreDtTm,
		NumberOfTxs:             grpHdr.NbOfTxs,
		SettlementMethod:        grpHdr.SttlmInf.SttlmMtd,
		ControlSum:              grpHdr.CtrlSum,
		SettlementDate:          txn.IntrBkSttlmDt,
		SettlementPriority:      txn.SttlmPrty,
		RemittanceInfo:          txn.RmtInf.Ustrd,
		Purpose:                 txn.Purp.Cd,
		ChargeBearer:            txn.ChrgBr,
		Charges:                 charges,
		// ReferenceNumber :
		// TotalTaxAmount
	}

	// x, _ := json.MarshalIndent(po, "", "  ")
	// log.Println("mapping po: ", string(x))
	return nil, &po
}

func mapAmount(amount isomodels.Amount) *models.Amount {
	return &models.Amount{
		Value:    amount.Value,
		Currency: amount.Currency,
	}
}
func mapAgent(agent isomodels.Agent) *models.Agent {
	return &models.Agent{
		FiiBicfi:                       agent.FinInstnId.BICFI,
		FiiClearingSystemIdCode:        agent.FinInstnId.ClrSysMmbId.ClrSysId.Cd,
		FiiClearingSystemIdProprietary: agent.FinInstnId.ClrSysMmbId.ClrSysId.Prtry,
		FiiMemberId:                    agent.FinInstnId.ClrSysMmbId.MmbId,
		FiiLei:                         agent.FinInstnId.LEI,
		FiiName:                        agent.FinInstnId.Nm,
		FiiPostalAddress:               *mapPostalAddress(agent.FinInstnId.PstlAdr),
		FiiOtherId:                     agent.FinInstnId.Othr.Id,
		FiiOtherIssuer:                 agent.FinInstnId.Othr.Issr,
		FiiOtherSchemeNameCode:         agent.FinInstnId.Othr.SchmeNm.Cd,
		FiiOtherSchemeNameProprietary:  agent.FinInstnId.Othr.SchmeNm.Prtry,
		BiId:                           agent.BrnchId.Id,
		BiLei:                          agent.BrnchId.LEI,
		BiName:                         agent.BrnchId.Nm,
		BiPostalAddress:                *mapPostalAddress(agent.BrnchId.PstlAdr),
	}
}

func mapParty(party isomodels.Party) *models.Party {
	return &models.Party{
		Name:                     party.Nm,
		PostalAddress:            *mapPostalAddress(party.PstlAdr),
		OrgIdAnyBic:              party.Id.OrgId.AnyBIC,
		OrgIdLei:                 party.Id.OrgId.LEI,
		OrgIdOther:               *mapOrgIdOther(party.Id.OrgId.Othr),
		PrivateIdBirthDate:       party.Id.PrvtId.DtAndPlcOfBirth.BirthDt,
		PrivateIdCityOfBirth:     party.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth,
		PrivateIdCountryOfBirth:  party.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth,
		PrivateIdProvinceOfBirth: party.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth,
		PrivateIdOther:           *mapPrvtIdOther(party.Id.PrvtId.Othr),
		CountryOfResidence:       party.CtryOfRes,
		ContactDetails:           *mapContactDetails(party.CtctDtls),
	}
}

func mapAccount(account isomodels.Account) *models.Account {
	var acct = models.Account{
		Iban:            account.Id.IBAN,
		TypeCode:        account.Tp.Cd,
		TypeProprietary: account.Tp.Prtry,
		Currency:        account.Ccy,
		Name:            account.Nm,
		ProxyType:       account.Prxy.Tp,
		ProxyId:         account.Prxy.Id,
	}

	if account.Id.Othr != nil {
		if account.Id.Othr.Id != "" {
			acct.OtherId = account.Id.Othr.Id
		}
		if account.Id.Othr.Issr != "" {
			acct.OtherIssuer = account.Id.Othr.Issr
		}
		if account.Id.Othr.SchmeNm != nil {
			if account.Id.Othr.SchmeNm.Cd != "" {
				acct.OtherSchemeNameCode = account.Id.Othr.SchmeNm.Cd
			}
			if account.Id.Othr.SchmeNm.Prtry != "" {
				acct.OtherSchemeNameProprietary = account.Id.Othr.SchmeNm.Prtry
			}
		}
	}

	return &acct
}

func mapOrgIdOther(orgIdOthr []isomodels.OrgIdOthr) *[]models.OrgIdOthr {
	var orgIdOthers []models.OrgIdOthr

	for _, org := range orgIdOthr {
		orgIdOthers = append(orgIdOthers, models.OrgIdOthr{
			Id:                    org.Id,
			Issuer:                org.Issr,
			SchemeNameCode:        org.SchmeNm.Cd,
			SchemeNameProprietary: org.SchmeNm.Prtry,
		})
	}
	return &orgIdOthers
}

func mapPrvtIdOther(prvtIdOthr []isomodels.PrvtIdOthr) *[]models.PrivateIdOthr {
	var privateIdOthr []models.PrivateIdOthr

	for _, org := range prvtIdOthr {
		privateIdOthr = append(privateIdOthr, models.PrivateIdOthr{
			Id:                    org.Id,
			Issuer:                org.Issr,
			SchemeNameCode:        org.SchmeNm.Cd,
			SchemeNameProprietary: org.SchmeNm.Prtry,
		})
	}
	return &privateIdOthr
}

func mapContactDetails(ctctDtls isomodels.CtctDtls) *models.ContactDetails {
	return &models.ContactDetails{
		NamePrefix:      ctctDtls.NmPrfx,
		Name:            ctctDtls.Nm,
		PhoneNumber:     ctctDtls.PhneNb,
		MobileNumber:    ctctDtls.MobNb,
		FaxNumber:       ctctDtls.FaxNb,
		UrlAddress:      ctctDtls.URLAdr,
		EmailAddress:    ctctDtls.EmailAdr,
		EmailPurpose:    ctctDtls.EmailPurp,
		JobTitle:        ctctDtls.JobTitl,
		Responsibility:  ctctDtls.Rspnsblty,
		Department:      ctctDtls.Dept,
		PreferredMethod: ctctDtls.PrefrdMtd,
		Other:           *mapContactOther(ctctDtls.Othr),
	}
}

func mapContactOther(contactOthr []isomodels.ContactOthr) *[]models.ContactOthr {
	var contactOthers []models.ContactOthr

	for _, contact := range contactOthr {
		contactOthers = append(contactOthers, models.ContactOthr{
			Id:          contact.Id,
			ChannelType: contact.ChanlTp,
		})
	}
	return &contactOthers
}

func mapPostalAddress(pstlAdr isomodels.PstlAdr) *models.PostalAddress {
	var address = models.PostalAddress{
		AddressTypeCode:    pstlAdr.AdrTp.Cd,
		CareOf:             pstlAdr.CareOf,
		Department:         pstlAdr.Dept,
		SubDepartment:      pstlAdr.SubDept,
		StreetName:         pstlAdr.StrtNm,
		BuildingNumber:     pstlAdr.BldgNb,
		BuildingName:       pstlAdr.BldgNm,
		Floor:              pstlAdr.Flr,
		UnitNumber:         pstlAdr.UnitNb,
		PostBox:            pstlAdr.PstBx,
		Room:               pstlAdr.Room,
		PostalCode:         pstlAdr.PstCd,
		TownName:           pstlAdr.TwnNm,
		TownLocationName:   pstlAdr.TwnLctnNm,
		DistrictName:       pstlAdr.DstrctNm,
		CountrySubdivision: pstlAdr.CtrySubDvsn,
		Country:            pstlAdr.Ctry,
		AddressLine:        pstlAdr.AdrLine,
	}

	if pstlAdr.AdrTp.Prtry != nil {
		if pstlAdr.AdrTp.Prtry.Id != "" {
			address.AddressTypeProprietaryId = pstlAdr.AdrTp.Prtry.Id
		}
		if pstlAdr.AdrTp.Prtry.Issr != "" {
			address.AddressTypeProprietaryIssuer = pstlAdr.AdrTp.Prtry.Issr
		}
		if pstlAdr.AdrTp.Prtry.SchmeNm != "" {
			address.AddressTypeProprietarySchemeName = pstlAdr.AdrTp.Prtry.SchmeNm
		}
	}

	return &address
}
