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
			Amount: models.Amount{
				Value:    ch.Amt.Value,
				Currency: ch.Amt.Currency,
			},
			Agent: models.Agent{

				// map fields from ch.Agt

			},
			Type: models.ChrgsTp{
				Code:              ch.Tp.Cd,
				ProprietaryId:     ch.Tp.Prtry.Id,
				ProprietaryIssuer: ch.Tp.Prtry.Issr,
			},
		})
	}

	po = models.PaymentOrder{
		Debtor: models.Party{
			Name: txn.Dbtr.Nm,
			PostalAddress: models.PstlAdr{
				AddressTypeCode:    txn.Dbtr.PstlAdr.AdrTp.Cd,
				CareOf:             txn.Dbtr.PstlAdr.CareOf,
				Department:         txn.Dbtr.PstlAdr.Dept,
				SubDepartment:      txn.Dbtr.PstlAdr.SubDept,
				StreetName:         txn.Dbtr.PstlAdr.StrtNm,
				BuildingNumber:     txn.Dbtr.PstlAdr.BldgNb,
				BuildingName:       txn.Dbtr.PstlAdr.BldgNm,
				Floor:              txn.Dbtr.PstlAdr.Flr,
				UnitNumber:         txn.Dbtr.PstlAdr.UnitNb,
				PostBox:            txn.Dbtr.PstlAdr.PstBx,
				Room:               txn.Dbtr.PstlAdr.Room,
				PostalCode:         txn.Dbtr.PstlAdr.PstCd,
				TownName:           txn.Dbtr.PstlAdr.TwnNm,
				TownLocationName:   txn.Dbtr.PstlAdr.TwnLctnNm,
				DistrictName:       txn.Dbtr.PstlAdr.DstrctNm,
				CountrySubdivision: txn.Dbtr.PstlAdr.CtrySubDvsn,
				Country:            txn.Dbtr.PstlAdr.Ctry,
				AddressLine:        txn.Dbtr.PstlAdr.AdrLine,
			},
			OrgIdAnyBic:              txn.Dbtr.Id.OrgId.AnyBIC,
			OrgIdLei:                 txn.Dbtr.Id.OrgId.LEI,
			OrgIdOther:               *mapOrgIdOther(txn.Dbtr.Id.OrgId.Othr),
			PrivateIdBirthDate:       txn.Dbtr.Id.PrvtId.DtAndPlcOfBirth.BirthDt,
			PrivateIdCityOfBirth:     txn.Dbtr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth,
			PrivateIdCountryOfBirth:  txn.Dbtr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth,
			PrivateIdProvinceOfBirth: txn.Dbtr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth,
			PrivateIdOther:           *mapPrvtIdOther(txn.Dbtr.Id.PrvtId.Othr),

			CountryOfResidence: txn.Dbtr.CtryOfRes,
			ContactDetails: models.CtctDtls{
				NamePrefix:      txn.Dbtr.CtctDtls.NmPrfx,
				Name:            txn.Dbtr.CtctDtls.Nm,
				PhoneNumber:     txn.Dbtr.CtctDtls.PhneNb,
				MobileNumber:    txn.Dbtr.CtctDtls.MobNb,
				FaxNumber:       txn.Dbtr.CtctDtls.FaxNb,
				UrlAddress:      txn.Dbtr.CtctDtls.URLAdr,
				EmailAddress:    txn.Dbtr.CtctDtls.EmailAdr,
				EmailPurpose:    txn.Dbtr.CtctDtls.EmailPurp,
				JobTitle:        txn.Dbtr.CtctDtls.JobTitl,
				Responsibility:  txn.Dbtr.CtctDtls.Rspnsblty,
				Department:      txn.Dbtr.CtctDtls.Dept,
				Other:           *mapContactOther(txn.Dbtr.CtctDtls.Othr),
				PreferredMethod: txn.Dbtr.CtctDtls.PrefrdMtd,
			},
		},
		DebtorAcct: models.Account{
			Iban:            txn.DbtrAcct.Id.IBAN,
			TypeCode:        txn.DbtrAcct.Tp.Cd,
			TypeProprietary: txn.DbtrAcct.Tp.Prtry,
			Currency:        txn.DbtrAcct.Ccy,
			Name:            txn.DbtrAcct.Nm,
			ProxyType:       txn.DbtrAcct.Prxy.Tp,
			ProxyId:         txn.DbtrAcct.Prxy.Id,
		},

		Creditor: models.Party{
			Name: txn.Cdtr.Nm,
			PostalAddress: models.PstlAdr{
				AddressTypeCode:                  txn.Cdtr.PstlAdr.AdrTp.Cd,
				AddressTypeProprietaryId:         txn.Cdtr.PstlAdr.AdrTp.Prtry.Id,
				AddressTypeProprietaryIssuer:     txn.Cdtr.PstlAdr.AdrTp.Prtry.Id,
				AddressTypeProprietarySchemeName: txn.Cdtr.PstlAdr.AdrTp.Prtry.SchmeNm,
				CareOf:                           txn.Cdtr.PstlAdr.CareOf,
				Department:                       txn.Cdtr.PstlAdr.Dept,
				SubDepartment:                    txn.Cdtr.PstlAdr.SubDept,
				StreetName:                       txn.Cdtr.PstlAdr.StrtNm,
				BuildingNumber:                   txn.Cdtr.PstlAdr.BldgNb,
				BuildingName:                     txn.Cdtr.PstlAdr.BldgNm,
				Floor:                            txn.Cdtr.PstlAdr.Flr,
				UnitNumber:                       txn.Cdtr.PstlAdr.UnitNb,
				PostBox:                          txn.Cdtr.PstlAdr.PstBx,
				Room:                             txn.Cdtr.PstlAdr.Room,
				PostalCode:                       txn.Cdtr.PstlAdr.PstCd,
				TownName:                         txn.Cdtr.PstlAdr.TwnNm,
				TownLocationName:                 txn.Cdtr.PstlAdr.TwnLctnNm,
				DistrictName:                     txn.Cdtr.PstlAdr.DstrctNm,
				CountrySubdivision:               txn.Cdtr.PstlAdr.CtrySubDvsn,
				Country:                          txn.Cdtr.PstlAdr.Ctry,
				AddressLine:                      txn.Cdtr.PstlAdr.AdrLine,
			},
			OrgIdAnyBic:              txn.Cdtr.Id.OrgId.AnyBIC,
			OrgIdLei:                 txn.Cdtr.Id.OrgId.LEI,
			OrgIdOther:               *mapOrgIdOther(txn.Cdtr.Id.OrgId.Othr),
			PrivateIdBirthDate:       txn.Cdtr.Id.PrvtId.DtAndPlcOfBirth.BirthDt,
			PrivateIdCityOfBirth:     txn.Cdtr.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth,
			PrivateIdCountryOfBirth:  txn.Cdtr.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth,
			PrivateIdProvinceOfBirth: txn.Cdtr.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth,
			PrivateIdOther:           *mapPrvtIdOther(txn.Cdtr.Id.PrvtId.Othr),
			CountryOfResidence:       txn.Cdtr.CtryOfRes,
			ContactDetails: models.CtctDtls{
				NamePrefix:      txn.Cdtr.CtctDtls.NmPrfx,
				Name:            txn.Cdtr.CtctDtls.Nm,
				PhoneNumber:     txn.Cdtr.CtctDtls.PhneNb,
				MobileNumber:    txn.Cdtr.CtctDtls.MobNb,
				FaxNumber:       txn.Cdtr.CtctDtls.FaxNb,
				UrlAddress:      txn.Cdtr.CtctDtls.URLAdr,
				EmailAddress:    txn.Cdtr.CtctDtls.EmailAdr,
				EmailPurpose:    txn.Cdtr.CtctDtls.EmailPurp,
				JobTitle:        txn.Cdtr.CtctDtls.JobTitl,
				Responsibility:  txn.Cdtr.CtctDtls.Rspnsblty,
				Department:      txn.Cdtr.CtctDtls.Dept,
				Other:           *mapContactOther(txn.Cdtr.CtctDtls.Othr),
				PreferredMethod: txn.Cdtr.CtctDtls.PrefrdMtd,
			},
		},

		CreditorAcct: models.Account{
			Iban:                       txn.CdtrAcct.Id.IBAN,
			TypeCode:                   txn.CdtrAcct.Tp.Cd,
			TypeProprietary:            txn.CdtrAcct.Tp.Prtry,
			Currency:                   txn.CdtrAcct.Ccy,
			Name:                       txn.CdtrAcct.Nm,
			ProxyType:                  txn.CdtrAcct.Prxy.Tp,
			ProxyId:                    txn.CdtrAcct.Prxy.Id,
			OtherId:                    txn.CdtrAcct.Id.Othr.Id,
			OtherIssuer:                txn.CdtrAcct.Id.Othr.Issr,
			OtherSchemeNameCode:        txn.CdtrAcct.Id.Othr.SchmeNm.Cd,
			OtherSchemeNameProprietary: txn.CdtrAcct.Id.Othr.SchmeNm.Prtry,
		},

		SettlementAmount: models.Amount{
			Value:    txn.IntrBkSttlmAmt.Value,
			Currency: txn.IntrBkSttlmAmt.Currency,
		},

		InstructionId:           txn.PmtId.InstrId,
		TransactionId:           txn.PmtId.TxId,
		EndToEndId:              txn.PmtId.EndToEndId,
		UETR:                    txn.PmtId.UETR,
		ClearingSystemReference: txn.PmtId.ClrSysRef,

		MsgId: grpHdr.MsgId,

		CreationDateTime: grpHdr.CreDtTm,
		NumberOfTxs:      grpHdr.NbOfTxs,
		SettlementMethod: grpHdr.SttlmInf.SttlmMtd,
		ControlSum:       grpHdr.CtrlSum,

		SettlementDate:     txn.IntrBkSttlmDt,
		SettlementPriority: txn.SttlmPrty,
		RemittanceInfo:     txn.RmtInf.Ustrd,
		Purpose:            txn.Purp.Cd,

		ChargeBearer: txn.ChrgBr,
		Charges:      charges,
		// ReferenceNumber :
		// TotalTaxAmount
	}

	if txn.Dbtr.PstlAdr.AdrTp.Prtry != nil {
		if txn.Dbtr.PstlAdr.AdrTp.Prtry.Id != "" {
			po.Debtor.PostalAddress.AddressTypeProprietaryId = txn.Dbtr.PstlAdr.AdrTp.Prtry.Id
		}
		if txn.Dbtr.PstlAdr.AdrTp.Prtry.Issr != "" {
			po.Debtor.PostalAddress.AddressTypeProprietaryIssuer = txn.Dbtr.PstlAdr.AdrTp.Prtry.Issr
		}
		if txn.Dbtr.PstlAdr.AdrTp.Prtry.SchmeNm != "" {
			po.Debtor.PostalAddress.AddressTypeProprietarySchemeName = txn.Dbtr.PstlAdr.AdrTp.Prtry.SchmeNm
		}
	}

	if txn.Cdtr.PstlAdr.AdrTp.Prtry != nil {
		if txn.Cdtr.PstlAdr.AdrTp.Prtry.Id != "" {
			po.Creditor.PostalAddress.AddressTypeProprietaryId = txn.Cdtr.PstlAdr.AdrTp.Prtry.Id
		}
		if txn.Cdtr.PstlAdr.AdrTp.Prtry.Issr != "" {
			po.Creditor.PostalAddress.AddressTypeProprietaryIssuer = txn.Cdtr.PstlAdr.AdrTp.Prtry.Issr
		}
		if txn.Cdtr.PstlAdr.AdrTp.Prtry.SchmeNm != "" {
			po.Creditor.PostalAddress.AddressTypeProprietarySchemeName = txn.Cdtr.PstlAdr.AdrTp.Prtry.SchmeNm
		}
	}

	// x, _ := json.MarshalIndent(po, "", "  ")
	// log.Println("mapping po: ", string(x))
	return nil, &po
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

func mapPrvtIdOther(orgIdOthr []isomodels.PrvtIdOthr) *[]models.PrivateIdOthr {
	var orgIdOthers []models.PrivateIdOthr

	for _, org := range orgIdOthr {
		orgIdOthers = append(orgIdOthers, models.PrivateIdOthr{
			Id:                    org.Id,
			Issuer:                org.Issr,
			SchemeNameCode:        org.SchmeNm.Cd,
			SchemeNameProprietary: org.SchmeNm.Prtry,
		})
	}
	return &orgIdOthers
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
