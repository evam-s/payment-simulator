package mapping

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"payment-simulator/internal/iso20022/isomodels"
	"payment-simulator/internal/models"

	"github.com/matoous/go-nanoid/v2"
)

func MapToPacs002(pos []*models.PaymentOrder, status map[string]string) (*isomodels.Pacs002, map[string]string) {
	if len(pos) < 1 {
		log.Println("Payment Order data must be sent to Generate PACS002.")
		return nil, nil
	}

	stsIdMap := make(map[string]string) // maps are reference datatypes. you dont make them with pointers. they themselves can be passed as pointers, and their underlying data will change if edited in another func.
	pacs002 := isomodels.Pacs002{
		FIToFIPmtStsRpt: &isomodels.FIToFIPmtStsRpt{
			GrpHdr:            &isomodels.GrpHdrPacs002{},
			OrgnlGrpInfAndSts: mapOrgnlGrpInfAndSts(pos),
			TxInfAndSts:       mapTxInfAndSts(pos, status, stsIdMap),
		}}

	return &pacs002, stsIdMap
}

func mapOrgnlGrpInfAndSts(pos []*models.PaymentOrder) []*isomodels.OrgnlGrpInfAndSts {
	var orgnlGrpInfAndSts []*isomodels.OrgnlGrpInfAndSts
	for _, po := range pos {
		if po != nil {
			orgnlGrpInfAndSts = append(orgnlGrpInfAndSts, &isomodels.OrgnlGrpInfAndSts{
				OrgnlMsgId:   po.MessageId,
				OrgnlMsgNmId: po.MessageNameId,
				OrgnlCreDtTm: po.CreationDateTime,
				OrgnlNbOfTxs: po.NumberOfTxs,
				OrgnlCtrlSum: po.ControlSum,
			})
		}
	}

	return orgnlGrpInfAndSts
}

func mapTxInfAndSts(pos []*models.PaymentOrder, status map[string]string, stsIdMap map[string]string) []*isomodels.TxInfAndSts {
	var txInfAndSts []*isomodels.TxInfAndSts
	for _, po := range pos {
		if po != nil {
			settlmKey := sha256.Sum256([]byte(po.InstructionId + po.EndToEndId + po.TransactionId + po.UETR))
			stsId, _ := gonanoid.New(12)
			stsIdMap[po.TransactionId] = stsId

			txSts := isomodels.TxInfAndSts{
				StsId:           stsId,
				OrgnlInstrId:    po.InstructionId,
				OrgnlEndToEndId: po.EndToEndId,
				OrgnlTxId:       po.TransactionId,
				OrgnlUETR:       po.UETR,
				TxSts:           status[po.TransactionId],
				ChrgsInf:        mapChrgsInf(po.Charges),
				AccptncDtTm:     po.TxnAcceptanceDateTime,
				AcctSvcrRef:     po.Id,
				ClrSysRef:       po.ClearingSystemReference,
				CdtSttlmKey:     hex.EncodeToString(settlmKey[:]),
				OrgnlTxRef:      mapOrgnlTxRef(po),
				SplmtryData:     mapSplmtryData(po.SupplementaryData),
			}

			if po.MessageId != "" || po.MessageNameId != "" || po.CreationDateTime != "" {
				txSts.OrgnlGrpInf = &isomodels.OrgnlGrpInf{
					OrgnlMsgId:   po.MessageId,
					OrgnlMsgNmId: po.MessageNameId,
					OrgnlCreDtTm: po.CreationDateTime,
				}
			}

			if po.ProcessingDateTime != "" {
				txSts.PrcgDt = &isomodels.DtAndDtTmChoice{
					DtTm: po.ProcessingDateTime,
				}
			}

			if po.EffectiveSettlementDateTime != "" {
				txSts.FctvIntrBkSttlmDt = &isomodels.DtAndDtTmChoice{
					DtTm: po.EffectiveSettlementDateTime,
				}
			}

			txInfAndSts = append(txInfAndSts, &txSts)
		}
	}

	return txInfAndSts
}

func mapSplmtryData(supplementaryData []*models.SupplementaryData) []*isomodels.SplmtryData {
	var supplData []*isomodels.SplmtryData
	for _, splData := range supplementaryData {
		if splData != nil {
			splD := isomodels.SplmtryData{
				PlcAndNm: splData.PlaceAndName,
			}

			if splData.Envelope != "" {
				splD.Envlp = &isomodels.Envlp{
					Data: splData.Envelope,
				}
			}

			supplData = append(supplData, &splD)
		}
	}

	return supplData
}

func mapChrgsInf(charges []*models.Charges) []*isomodels.ChrgsInf {
	if len(charges) < 1 {
		return nil
	}

	var chrgsInf []*isomodels.ChrgsInf
	for _, chrgs := range charges {
		if chrgs != nil {
			chg := isomodels.ChrgsInf{
				Amt: mapPoAmount(chrgs.Amount),
				Agt: mapPoAgent(chrgs.Agent),
			}

			if chrgs.Type != nil {
				chg.Tp = &isomodels.ChrgsTp{
					Cd: chrgs.Type.Code,
				}

				if chrgs.Type.ProprietaryId != "" || chrgs.Type.ProprietaryIssuer != "" {
					chg.Tp.Prtry = &isomodels.ChrgsTpPrtry{
						Id:   chrgs.Type.ProprietaryId,
						Issr: chrgs.Type.ProprietaryIssuer,
					}
				}
			}

			chrgsInf = append(chrgsInf, &chg)
		}
	}

	return chrgsInf
}

func mapPoAmount(amount *models.Amount) *isomodels.Amount {
	if amount == nil {
		return nil
	}

	return &isomodels.Amount{
		Value:    amount.Value,
		Currency: amount.Currency,
	}
}

func mapPoAgent(agent *models.Agent) *isomodels.Agent {
	if agent == nil {
		return nil
	}

	agt := isomodels.Agent{}

	if agent.FiiBicfi != "" || agent.FiiClearingSystemIdCode != "" || agent.FiiClearingSystemIdProprietary != "" || agent.FiiMemberId != "" || agent.FiiLei != "" || agent.FiiName != "" {
		agt.FinInstnId = &isomodels.FinInstnId{
			BICFI: agent.FiiBicfi,
			LEI:   agent.FiiLei,
			Nm:    agent.FiiName,
		}

		if agent.FiiClearingSystemIdCode != "" || agent.FiiClearingSystemIdProprietary != "" || agent.FiiMemberId != "" {
			agt.FinInstnId.ClrSysMmbId = &isomodels.ClrSysMmbId{
				MmbId: agent.FiiMemberId,
			}

			if agent.FiiClearingSystemIdCode != "" || agent.FiiClearingSystemIdProprietary != "" {
				agt.FinInstnId.ClrSysMmbId.ClrSysId = &isomodels.ClrSysId{
					Cd:    agent.FiiClearingSystemIdCode,
					Prtry: agent.FiiClearingSystemIdProprietary,
				}
			}
		}

		if agent.FiiOtherId != "" || agent.FiiOtherIssuer != "" || agent.FiiOtherSchemeNameCode != "" || agent.FiiOtherSchemeNameProprietary != "" {
			agt.FinInstnId.Othr = &isomodels.FinInstnOthr{
				Id:   agent.FiiOtherId,
				Issr: agent.FiiOtherIssuer,
			}

			if agent.FiiOtherSchemeNameCode != "" || agent.FiiOtherSchemeNameProprietary != "" {
				agt.FinInstnId.Othr.SchmeNm = &isomodels.FinInstnSchmeNm{
					Cd:    agent.FiiOtherSchemeNameCode,
					Prtry: agent.FiiOtherSchemeNameProprietary,
				}
			}
		}
	}

	if agent.BiId != "" || agent.BiLei != "" || agent.BiName != "" {
		agt.BrnchId = &isomodels.BrnchId{
			Id:  agent.BiId,
			LEI: agent.BiLei,
			Nm:  agent.BiName,
		}
	}

	return &agt
}

func mapPstlAdr(postalAddress *models.PostalAddress) *isomodels.PstlAdr {
	if postalAddress == nil {
		return nil
	}

	adr := isomodels.PstlAdr{
		CareOf:      postalAddress.CareOf,
		Dept:        postalAddress.Department,
		SubDept:     postalAddress.SubDepartment,
		StrtNm:      postalAddress.StreetName,
		BldgNb:      postalAddress.BuildingNumber,
		BldgNm:      postalAddress.BuildingName,
		Flr:         postalAddress.Floor,
		UnitNb:      postalAddress.UnitNumber,
		PstBx:       postalAddress.PostBox,
		Room:        postalAddress.Room,
		PstCd:       postalAddress.PostalCode,
		TwnNm:       postalAddress.TownName,
		TwnLctnNm:   postalAddress.TownLocationName,
		DstrctNm:    postalAddress.DistrictName,
		CtrySubDvsn: postalAddress.CountrySubdivision,
		Ctry:        postalAddress.Country,
		AdrLine:     postalAddress.AddressLine,
	}

	if postalAddress.AddressTypeCode != "" || postalAddress.AddressTypeProprietaryId != "" || postalAddress.AddressTypeProprietaryIssuer != "" || postalAddress.AddressTypeProprietarySchemeName != "" {
		adr.AdrTp = &isomodels.AdrTp{
			Cd: postalAddress.AddressTypeCode,
		}

		if postalAddress.AddressTypeProprietaryId != "" || postalAddress.AddressTypeProprietaryIssuer != "" || postalAddress.AddressTypeProprietarySchemeName != "" {
			adr.AdrTp.Prtry = &isomodels.AdrTpPrtry{
				Id:      postalAddress.AddressTypeProprietaryId,
				Issr:    postalAddress.AddressTypeProprietaryIssuer,
				SchmeNm: postalAddress.AddressTypeProprietarySchemeName,
			}
		}
	}

	return &adr
}

func mapOrgnlTxRef(po *models.PaymentOrder) *isomodels.OrgnlTxRefPacs002 {
	if po == nil {
		return nil
	}

	orgnl := isomodels.OrgnlTxRefPacs002{
		IntrBkSttlmDt: po.SettlementDate,
		PmtTpInf:      mapPmtTpInf(po.PaymentTypeInfo),
		PmtMtd:        "TRF",
		Dbtr: &isomodels.PartyOrAgentChoice{
			Pty: mapPoParty(po.Debtor),
		},
		DbtrAcct:    mapPoAccount(po.DebtorAcct),
		DbtrAgt:     mapPoAgent(po.DebtorAgent),
		DbtrAgtAcct: mapPoAccount(po.DebtorAgentAcct),
		CdtrAgt:     mapPoAgent(po.CreditorAgent),
		CdtrAgtAcct: mapPoAccount(po.CreditorAgentAcct),
		Cdtr: &isomodels.PartyOrAgentChoice{
			Pty: mapPoParty(po.Creditor),
		},
		CdtrAcct: mapPoAccount(po.CreditorAcct),
	}

	if po.SettlementMethod != "" || po.SettlementAcct != nil {
		orgnl.SttlmInf = &isomodels.SttlmInf{
			SttlmMtd:  po.SettlementMethod,
			SttlmAcct: mapPoAccount(po.SettlementAcct),
		}
	}

	if len(po.UnstructuredRemittanceInfo) > 0 {
		for _, v := range po.UnstructuredRemittanceInfo {
			if v != "" {
				if orgnl.RmtInf == nil {
					orgnl.RmtInf = &isomodels.RmtInf{}
				}

				orgnl.RmtInf.Ustrd = append(orgnl.RmtInf.Ustrd, v)
			}
		}
	}

	if po.UltimateDebtor != nil {
		orgnl.UltmtDbtr = &isomodels.PartyOrAgentChoice{
			Pty: mapPoParty(po.UltimateDebtor),
		}
	}

	if po.UltimateCreditor != nil {
		orgnl.UltmtCdtr = &isomodels.PartyOrAgentChoice{
			Pty: mapPoParty(po.UltimateCreditor),
		}
	}

	if po.PurposeCode != "" || po.PurposeProprietary != "" {
		orgnl.Purp = &isomodels.Purp{
			Cd:    po.PurposeCode,
			Prtry: po.PurposeProprietary,
		}
	}

	if po.SettlementAmount != nil {
		orgnl.IntrBkSttlmAmt = &isomodels.Amount{
			Value:    po.SettlementAmount.Value,
			Currency: po.SettlementAmount.Currency,
		}
	}

	return &orgnl
}

func mapPoAccount(account *models.Account) *isomodels.Account {
	if account == nil {
		return nil
	}

	act := isomodels.Account{
		Ccy: account.Currency,
		Nm:  account.Name,
	}

	if account.Iban != "" || account.OtherId != "" || account.OtherIssuer != "" || account.OtherSchemeNameCode != "" || account.OtherSchemeNameProprietary != "" {
		act.Id = &isomodels.AccountId{
			IBAN: account.Iban,
		}

		if account.OtherId != "" || account.OtherIssuer != "" || account.OtherSchemeNameCode != "" || account.OtherSchemeNameProprietary != "" {
			act.Id.Othr = &isomodels.AccountOthr{
				Id:   account.OtherId,
				Issr: account.OtherIssuer,
			}

			if account.OtherSchemeNameCode != "" || account.OtherSchemeNameProprietary != "" {
				act.Id.Othr.SchmeNm = &isomodels.AccountSchmeNm{
					Cd:    account.OtherSchemeNameCode,
					Prtry: account.OtherSchemeNameProprietary,
				}
			}
		}
	}

	if account.TypeCode != "" || account.TypeProprietary != "" {
		act.Tp = &isomodels.AccountType{
			Cd:    account.TypeCode,
			Prtry: account.TypeProprietary,
		}
	}

	if account.ProxyType != "" || account.ProxyId != "" {
		act.Prxy = &isomodels.AccountProxy{
			Tp: account.ProxyType,
			Id: account.ProxyId,
		}
	}

	return &act
}

func mapPmtTpInf(pti *models.PaymentTypeInfo) *isomodels.PmtTpInfPacs002 {
	if pti == nil {
		return nil
	}

	pmt := isomodels.PmtTpInfPacs002{
		InstrPrty: pti.InstructionPriority,
		ClrChanl:  pti.ClearingChannel,
		SvcLvl:    mapServiceLevel(pti.ServiceLevel),
	}

	if pti.LocalInstrumentCode != "" || pti.LocalInstrumentProprietary != "" {
		pmt.LclInstrm = &isomodels.LclInstrm{
			Cd:    pti.LocalInstrumentCode,
			Prtry: pti.LocalInstrumentProprietary,
		}
	}

	if pti.CategoryPurposeCode != "" || pti.CategoryPurposeProprietary != "" {
		pmt.CtgyPurp = &isomodels.CtgyPurp{
			Cd:    pti.CategoryPurposeCode,
			Prtry: pti.CategoryPurposeProprietary,
		}
	}

	return &pmt
}

func mapServiceLevel(serviceLevels []*models.CodeOrProprietary) []*isomodels.SvcLvl {
	if serviceLevels == nil {
		return nil
	}

	var result []*isomodels.SvcLvl
	for _, serviceLevel := range serviceLevels {
		if serviceLevel != nil {
			result = append(result, &isomodels.SvcLvl{
				Cd:    serviceLevel.Code,
				Prtry: serviceLevel.Proprietary,
			})
		}
	}

	return result
}

func mapPoParty(party *models.Party) *isomodels.Party {
	if party == nil {
		return nil
	}

	pty := isomodels.Party{
		Nm:        party.Name,
		PstlAdr:   mapPstlAdr(party.PostalAddress),
		CtryOfRes: party.CountryOfResidence,
		CtctDtls:  mapCtctDtls(party.ContactDetails),
	}

	if party.OrgIdAnyBic != "" || party.OrgIdLei != "" || len(party.OrgIdOther) > 0 || party.PrivateIdBirthDate != "" || party.PrivateIdCityOfBirth != "" || party.PrivateIdCountryOfBirth != "" || party.PrivateIdProvinceOfBirth != "" || len(party.PrivateIdOther) > 0 {
		pty.Id = &isomodels.PartyId{}

		if party.OrgIdAnyBic != "" || party.OrgIdLei != "" || len(party.OrgIdOther) > 0 {
			pty.Id.OrgId = &isomodels.OrgId{
				AnyBIC: party.OrgIdAnyBic,
				LEI:    party.OrgIdLei,
				Othr:   mapPoOrgIdOthr(party.OrgIdOther),
			}
		}

		if party.PrivateIdBirthDate != "" || party.PrivateIdCityOfBirth != "" || party.PrivateIdCountryOfBirth != "" || party.PrivateIdProvinceOfBirth != "" || len(party.PrivateIdOther) > 0 {
			pty.Id.PrvtId.Othr = mapPoPrvtIdOthr(party.PrivateIdOther)

			if party.PrivateIdBirthDate != "" || party.PrivateIdCityOfBirth != "" || party.PrivateIdCountryOfBirth != "" || party.PrivateIdProvinceOfBirth != "" {
				pty.Id.PrvtId.DtAndPlcOfBirth = &isomodels.DateAndPlaceOfBirth{
					BirthDt:     party.PrivateIdBirthDate,
					CityOfBirth: party.PrivateIdCityOfBirth,
					CtryOfBirth: party.PrivateIdCountryOfBirth,
					PrvcOfBirth: party.PrivateIdProvinceOfBirth,
				}
			}
		}
	}

	return &pty
}

func mapPoOrgIdOthr(orgIdOthers []*models.OrgIdOthr) []*isomodels.OrgIdOthr {
	if len(orgIdOthers) < 1 {
		return nil
	}

	var isomodelsOthers []*isomodels.OrgIdOthr
	for _, org := range orgIdOthers {
		if org != nil {
			oio := isomodels.OrgIdOthr{
				Id:   org.Id,
				Issr: org.Issuer,
			}

			if org.SchemeNameCode != "" || org.SchemeNameProprietary != "" {
				oio.SchmeNm = &isomodels.OrgSchmeNm{
					Cd:    org.SchemeNameCode,
					Prtry: org.SchemeNameProprietary,
				}
			}

			isomodelsOthers = append(isomodelsOthers, &oio)
		}
	}

	return isomodelsOthers
}

func mapPoPrvtIdOthr(privateIdOthers []*models.PrivateIdOthr) []*isomodels.PrvtIdOthr {
	if len(privateIdOthers) < 1 {
		return nil
	}

	var isoPrvtOthers []*isomodels.PrvtIdOthr
	for _, prvt := range privateIdOthers {
		if prvt != nil {
			pio := isomodels.PrvtIdOthr{
				Id:   prvt.Id,
				Issr: prvt.Issuer,
			}

			if prvt.SchemeNameCode != "" || prvt.SchemeNameProprietary != "" {
				pio.SchmeNm = &isomodels.PersonSchmeNm{
					Cd:    prvt.SchemeNameCode,
					Prtry: prvt.SchemeNameProprietary,
				}
			}

			isoPrvtOthers = append(isoPrvtOthers, &pio)
		}
	}

	return isoPrvtOthers
}

func mapCtctDtls(contact *models.ContactDetails) *isomodels.CtctDtls {
	if contact == nil {
		return nil
	}

	return &isomodels.CtctDtls{
		NmPrfx:    contact.NamePrefix,
		Nm:        contact.Name,
		PhneNb:    contact.PhoneNumber,
		MobNb:     contact.MobileNumber,
		FaxNb:     contact.FaxNumber,
		URLAdr:    contact.UrlAddress,
		EmailAdr:  contact.EmailAddress,
		EmailPurp: contact.EmailPurpose,
		JobTitl:   contact.JobTitle,
		Rspnsblty: contact.Responsibility,
		Dept:      contact.Department,
		PrefrdMtd: contact.PreferredMethod,
		Othr:      mapCtctDtlsOthr(contact.Other),
	}
}

func mapCtctDtlsOthr(contactOthers []*models.ContactOthr) []*isomodels.ContactOthr {
	if len(contactOthers) < 1 {
		return nil
	}

	var isoContacts []*isomodels.ContactOthr
	for _, c := range contactOthers {
		if c != nil {
			isoContacts = append(isoContacts, &isomodels.ContactOthr{
				Id:      c.Id,
				ChanlTp: c.ChannelType,
			})
		}
	}

	return isoContacts
}
