package mapping

import (
	"log"
	"payment-simulator/internal/iso20022/isomodels"
	"payment-simulator/internal/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func MapXmlPacs008(c *gin.Context) (*isomodels.Pacs008, error) {
	var isoPacs isomodels.Pacs008
	if err := c.ShouldBindXML(&isoPacs); err != nil {
		log.Println("Error in Binding to XML:", isoPacs, "Error:", err)
		return nil, err
	} else {
		return &isoPacs, nil
	}
}

func MapPacs008ToPo(isoPacs *isomodels.Pacs008) (error, *models.PaymentOrder) {
	var po models.PaymentOrder
	grpHdr := isoPacs.FIToFICstmrCdtTrf.GrpHdr
	txn := isoPacs.FIToFICstmrCdtTrf.CdtTrfTxInf[0]
	var charges []*models.Charges
	for _, ch := range txn.ChrgsInf {
		if ch != nil {
			chg := models.Charges{
				Amount: mapAmount(ch.Amt),
				Agent:  mapIsoAgent(ch.Agt),
			}
			if ch.Tp != nil {
				chg.Type = &models.ChrgsTp{}
				chg.Type.Code = ch.Tp.Cd
				if ch.Tp.Prtry != nil {
					chg.Type.ProprietaryId = ch.Tp.Prtry.Id
					chg.Type.ProprietaryIssuer = ch.Tp.Prtry.Issr
				}
			}
			charges = append(charges, &chg)
		}
	}

	po = models.PaymentOrder{
		Debtor:                  mapIsoParty(txn.Dbtr),
		UltimateDebtor:          mapIsoParty(txn.UltmtDbtr),
		DebtorAcct:              mapIsoAccount(txn.DbtrAcct),
		DebtorAgent:             mapIsoAgent(txn.DbtrAgt),
		DebtorAgentAcct:         mapIsoAccount(txn.DbtrAgtAcct),
		Creditor:                mapIsoParty(txn.Cdtr),
		UltimateCreditor:        mapIsoParty(txn.UltmtCdtr),
		CreditorAcct:            mapIsoAccount(txn.CdtrAcct),
		CreditorAgent:           mapIsoAgent(txn.CdtrAgt),
		CreditorAgentAcct:       mapIsoAccount(txn.CdtrAgtAcct),
		SettlementAmount:        mapAmount(txn.IntrBkSttlmAmt),
		MessageId:               grpHdr.MsgId,
		MessageNameId:           strings.Split(isoPacs.Xmlns, "xsd:")[1],
		CreationDateTime:        grpHdr.CreDtTm,
		NumberOfTxs:             grpHdr.NbOfTxs,
		ControlSum:              grpHdr.CtrlSum,
		SettlementDate:          txn.IntrBkSttlmDt,
		SettlementPriority:      txn.SttlmPrty,
		ChargeBearer:            txn.ChrgBr,
		Charges:                 charges,
		PaymentTypeInfo:         mapPaymentTypeInfo(txn.PmtTpInf),
		HeaderPaymentTypeInfo:   mapPaymentTypeInfo(grpHdr.PmtTpInf),
		SupplementaryData:       mapSupplementaryData(txn.SplmtryData),
		HeaderSupplementaryData: mapSupplementaryData(isoPacs.FIToFICstmrCdtTrf.SplmtryData),
		// ReferenceNumber :
		// TotalTaxAmount
	}

	if txn.PmtId != nil {
		po.UETR = txn.PmtId.UETR
		po.InstructionId = txn.PmtId.InstrId
		po.TransactionId = txn.PmtId.TxId
		po.EndToEndId = txn.PmtId.EndToEndId
		po.ClearingSystemReference = txn.PmtId.ClrSysRef
	}

	if grpHdr.SttlmInf != nil {
		po.SettlementMethod = grpHdr.SttlmInf.SttlmMtd
		po.SettlementAcct = mapIsoAccount(grpHdr.SttlmInf.SttlmAcct)
	}

	if txn.RmtInf != nil {
		po.UnstructuredRemittanceInfo = txn.RmtInf.Ustrd
	}

	if txn.Purp != nil {
		po.PurposeCode = txn.Purp.Cd
		po.PurposeProprietary = txn.Purp.Prtry
	}

	// x, _ := json.MarshalIndent(po, "", "  ")
	// log.Println("mapping po:", string(x))
	return nil, &po
}

func mapAmount(amount *isomodels.Amount) *models.Amount {
	if amount == nil {
		return nil
	}

	return &models.Amount{
		Value:    amount.Value,
		Currency: amount.Currency,
	}
}

func mapIsoAgent(agent *isomodels.Agent) *models.Agent {
	if agent == nil {
		return nil
	}

	agt := models.Agent{}

	if agent.FinInstnId != nil {
		agt.FiiLei = agent.FinInstnId.LEI
		agt.FiiName = agent.FinInstnId.Nm
		agt.FiiBicfi = agent.FinInstnId.BICFI
		agt.FiiPostalAddress = mapPostalAddress(agent.FinInstnId.PstlAdr)

		if agent.FinInstnId.ClrSysMmbId != nil {
			agt.FiiMemberId = agent.FinInstnId.ClrSysMmbId.MmbId

			if agent.FinInstnId.ClrSysMmbId.ClrSysId != nil {
				agt.FiiClearingSystemIdCode = agent.FinInstnId.ClrSysMmbId.ClrSysId.Cd
				agt.FiiClearingSystemIdProprietary = agent.FinInstnId.ClrSysMmbId.ClrSysId.Prtry
			}
		}

		if agent.FinInstnId.Othr != nil {
			agt.FiiOtherId = agent.FinInstnId.Othr.Id
			agt.FiiOtherIssuer = agent.FinInstnId.Othr.Issr

			if agent.FinInstnId.Othr.SchmeNm != nil {
				agt.FiiOtherSchemeNameCode = agent.FinInstnId.Othr.SchmeNm.Cd
				agt.FiiOtherSchemeNameProprietary = agent.FinInstnId.Othr.SchmeNm.Prtry
			}
		}
	}

	if agent.BrnchId != nil {
		agt.BiId = agent.BrnchId.Id
		agt.BiLei = agent.BrnchId.LEI
		agt.BiName = agent.BrnchId.Nm
		agt.BiPostalAddress = mapPostalAddress(agent.BrnchId.PstlAdr)
	}

	return &agt
}

func mapIsoParty(pty *isomodels.Party) *models.Party {
	if pty == nil {
		return nil
	}

	party := models.Party{
		Name:               pty.Nm,
		PostalAddress:      mapPostalAddress(pty.PstlAdr),
		CountryOfResidence: pty.CtryOfRes,
		ContactDetails:     mapContactDetails(pty.CtctDtls),
	}

	if pty.Id != nil {
		if pty.Id.OrgId != nil {
			party.OrgIdAnyBic = pty.Id.OrgId.AnyBIC
			party.OrgIdLei = pty.Id.OrgId.LEI
			party.OrgIdOther = mapOrgIdOther(pty.Id.OrgId.Othr)
		}

		if pty.Id.PrvtId != nil {
			if pty.Id.PrvtId.DtAndPlcOfBirth != nil {
				party.PrivateIdBirthDate = pty.Id.PrvtId.DtAndPlcOfBirth.BirthDt
				party.PrivateIdCityOfBirth = pty.Id.PrvtId.DtAndPlcOfBirth.CityOfBirth
				party.PrivateIdCountryOfBirth = pty.Id.PrvtId.DtAndPlcOfBirth.CtryOfBirth
				party.PrivateIdProvinceOfBirth = pty.Id.PrvtId.DtAndPlcOfBirth.PrvcOfBirth
			}
			party.PrivateIdOther = mapPrvtIdOther(pty.Id.PrvtId.Othr)
		}
	}

	return &party
}

func mapIsoAccount(account *isomodels.Account) *models.Account {
	if account == nil {
		return nil
	}

	acct := models.Account{
		Currency: account.Ccy,
		Name:     account.Nm,
	}

	if account.Id != nil {
		acct.Iban = account.Id.IBAN
	}

	if account.Tp != nil {
		acct.TypeCode = account.Tp.Cd
		acct.TypeProprietary = account.Tp.Prtry
	}

	if account.Prxy != nil {
		acct.ProxyId = account.Prxy.Id
		acct.ProxyType = account.Prxy.Tp
	}

	if account.Id != nil && account.Id.Othr != nil {
		acct.OtherId = account.Id.Othr.Id
		acct.OtherIssuer = account.Id.Othr.Issr
		if account.Id.Othr.SchmeNm != nil {
			acct.OtherSchemeNameCode = account.Id.Othr.SchmeNm.Cd
			acct.OtherSchemeNameProprietary = account.Id.Othr.SchmeNm.Prtry
		}
	}

	return &acct
}

func mapOrgIdOther(orgIdOthr []*isomodels.OrgIdOthr) []*models.OrgIdOthr {
	if len(orgIdOthr) < 1 {
		return nil
	}

	var orgIdOthers []*models.OrgIdOthr

	for _, org := range orgIdOthr {
		if org != nil {
			oio := &models.OrgIdOthr{
				Id:     org.Id,
				Issuer: org.Issr,
			}
			if org.SchmeNm != nil {
				oio.SchemeNameCode = org.SchmeNm.Cd
				oio.SchemeNameProprietary = org.SchmeNm.Prtry
			}
			orgIdOthers = append(orgIdOthers, oio)
		}
	}
	return orgIdOthers
}

func mapPrvtIdOther(prvtIdOthr []*isomodels.PrvtIdOthr) []*models.PrivateIdOthr {
	if len(prvtIdOthr) < 1 {
		return nil
	}

	var prvtIdOthers []*models.PrivateIdOthr

	for _, pvt := range prvtIdOthr {
		if pvt != nil {
			pio := &models.PrivateIdOthr{
				Id:     pvt.Id,
				Issuer: pvt.Issr,
			}
			if pvt.SchmeNm != nil {
				pio.SchemeNameCode = pvt.SchmeNm.Cd
				pio.SchemeNameProprietary = pvt.SchmeNm.Prtry
			}
			prvtIdOthers = append(prvtIdOthers, pio)
		}
	}
	return prvtIdOthers
}

func mapContactDetails(ctctDtls *isomodels.CtctDtls) *models.ContactDetails {
	if ctctDtls == nil {
		return nil
	}
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
		Other:           mapContactOther(ctctDtls.Othr),
	}
}

func mapContactOther(contactOthr []*isomodels.ContactOthr) []*models.ContactOthr {
	var contactOthers []*models.ContactOthr
	for _, contact := range contactOthr {
		if contact != nil {
			contactOthers = append(contactOthers, &models.ContactOthr{
				Id:          contact.Id,
				ChannelType: contact.ChanlTp,
			})
		}
	}
	return contactOthers
}

func mapPostalAddress(pstlAdr *isomodels.PstlAdr) *models.PostalAddress {
	if pstlAdr == nil {
		return nil
	}

	pa := models.PostalAddress{
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

	if pstlAdr.AdrTp != nil {
		pa.AddressTypeCode = pstlAdr.AdrTp.Cd
		if pstlAdr.AdrTp.Prtry != nil {
			pa.AddressTypeProprietaryId = pstlAdr.AdrTp.Prtry.Id
			pa.AddressTypeProprietaryIssuer = pstlAdr.AdrTp.Prtry.Issr
			pa.AddressTypeProprietarySchemeName = pstlAdr.AdrTp.Prtry.SchmeNm
		}
	}

	return &pa
}

func mapPaymentTypeInfo(pmtTpInf *isomodels.PmtTpInfPacs008) *models.PaymentTypeInfo {
	if pmtTpInf == nil {
		return nil
	}

	pmt := models.PaymentTypeInfo{
		InstructionPriority: pmtTpInf.InstrPrty,
		ClearingChannel:     pmtTpInf.ClrChanl,
		ServiceLevel:        mapSvcLvl(pmtTpInf.SvcLvl),
	}

	if pmtTpInf.LclInstrm != nil {
		pmt.LocalInstrumentCode = pmtTpInf.LclInstrm.Cd
		pmt.LocalInstrumentProprietary = pmtTpInf.LclInstrm.Prtry
	}
	if pmtTpInf.CtgyPurp != nil {
		pmt.CategoryPurposeCode = pmtTpInf.CtgyPurp.Cd
		pmt.CategoryPurposeProprietary = pmtTpInf.CtgyPurp.Prtry
	}

	return &pmt
}

func mapSvcLvl(serviceLevel []*isomodels.SvcLvl) []*models.CodeOrProprietary {
	var svcLvl []*models.CodeOrProprietary
	for _, sl := range serviceLevel {
		if sl != nil {
			svcLvl = append(svcLvl, &models.CodeOrProprietary{
				Code:        sl.Cd,
				Proprietary: sl.Prtry,
			})
		}
	}
	return svcLvl
}

func mapSupplementaryData(splmtryData []*isomodels.SplmtryData) []*models.SupplementaryData {
	var supplData []*models.SupplementaryData
	for _, splData := range splmtryData {
		if splData != nil {
			supplData = append(supplData, &models.SupplementaryData{
				PlaceAndName: splData.PlcAndNm,
				Envelope:     splData.Envlp.Data,
			})
		}
	}
	return supplData
}
