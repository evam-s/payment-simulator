package mapping

import (
	"github.com/gin-gonic/gin"
	"log"
	"payment-simulator/internal/iso20022/isomodels"
	"payment-simulator/internal/models"
	"strings"
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

	po = models.PaymentOrder{
		Debtor:                     mapIsoParty(txn.Dbtr),
		UltimateDebtor:             mapIsoParty(txn.UltmtDbtr),
		DebtorAcct:                 mapIsoAccount(txn.DbtrAcct),
		DebtorAgent:                mapIsoAgent(txn.DbtrAgt),
		DebtorAgentAcct:            mapIsoAccount(txn.DbtrAgtAcct),
		Creditor:                   mapIsoParty(txn.Cdtr),
		UltimateCreditor:           mapIsoParty(txn.UltmtCdtr),
		CreditorAcct:               mapIsoAccount(txn.CdtrAcct),
		CreditorAgent:              mapIsoAgent(txn.CdtrAgt),
		CreditorAgentAcct:          mapIsoAccount(txn.CdtrAgtAcct),
		SettlementAmount:           mapAmount(txn.IntrBkSttlmAmt),
		InstructionId:              txn.PmtId.InstrId,
		TransactionId:              txn.PmtId.TxId,
		EndToEndId:                 txn.PmtId.EndToEndId,
		UETR:                       txn.PmtId.UETR,
		ClearingSystemReference:    txn.PmtId.ClrSysRef,
		MessageId:                  grpHdr.MsgId,
		MessageNameId:              strings.Split(isoPacs.Xmlns, "xsd:")[1],
		CreationDateTime:           grpHdr.CreDtTm,
		NumberOfTxs:                grpHdr.NbOfTxs,
		SettlementMethod:           grpHdr.SttlmInf.SttlmMtd,
		SettlementAcct:             mapIsoAccount(grpHdr.SttlmInf.SttlmAcct),
		ControlSum:                 grpHdr.CtrlSum,
		SettlementDate:             txn.IntrBkSttlmDt,
		SettlementPriority:         txn.SttlmPrty,
		UnstructuredRemittanceInfo: txn.RmtInf.Ustrd,
		PurposeCode:                txn.Purp.Cd,
		PurposeProprietary:         txn.Purp.Prtry,
		ChargeBearer:               txn.ChrgBr,
		Charges:                    charges,
		PaymentTypeInfo:            mapPaymentTypeInfo(txn.PmtTpInf),
		HeaderPaymentTypeInfo:      mapPaymentTypeInfo(grpHdr.PmtTpInf),
		SupplementaryData:          mapSupplementaryData(txn.SplmtryData),
		HeaderSupplementaryData:    mapSupplementaryData(isoPacs.FIToFICstmrCdtTrf.SplmtryData),
		// ReferenceNumber :
		// TotalTaxAmount
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

	return &models.Agent{
		FiiBicfi:                       agent.FinInstnId.BICFI,
		FiiClearingSystemIdCode:        agent.FinInstnId.ClrSysMmbId.ClrSysId.Cd,
		FiiClearingSystemIdProprietary: agent.FinInstnId.ClrSysMmbId.ClrSysId.Prtry,
		FiiMemberId:                    agent.FinInstnId.ClrSysMmbId.MmbId,
		FiiLei:                         agent.FinInstnId.LEI,
		FiiName:                        agent.FinInstnId.Nm,
		FiiPostalAddress:               mapPostalAddress(agent.FinInstnId.PstlAdr),
		FiiOtherId:                     agent.FinInstnId.Othr.Id,
		FiiOtherIssuer:                 agent.FinInstnId.Othr.Issr,
		FiiOtherSchemeNameCode:         agent.FinInstnId.Othr.SchmeNm.Cd,
		FiiOtherSchemeNameProprietary:  agent.FinInstnId.Othr.SchmeNm.Prtry,
		BiId:                           agent.BrnchId.Id,
		BiLei:                          agent.BrnchId.LEI,
		BiName:                         agent.BrnchId.Nm,
		BiPostalAddress:                mapPostalAddress(agent.BrnchId.PstlAdr),
	}
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

	return &party
}

func mapIsoAccount(account *isomodels.Account) *models.Account {
	if account == nil {
		return nil
	}
	acct := models.Account{
		Iban:            account.Id.IBAN,
		TypeCode:        account.Tp.Cd,
		TypeProprietary: account.Tp.Prtry,
		Currency:        account.Ccy,
		Name:            account.Nm,
		ProxyType:       account.Prxy.Tp,
		ProxyId:         account.Prxy.Id,
	}

	if account.Id.Othr != nil {
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
