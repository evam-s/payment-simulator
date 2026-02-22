package mapping

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/matoous/go-nanoid/v2"
	"payment-simulator/internal/iso20022/isomodels"
	"payment-simulator/internal/models"
)

func MapToPacs002(pos []*models.PaymentOrder, status map[string]string) *isomodels.Pacs002 {
	var pacs002 isomodels.Pacs002
	pacs002.FIToFIPmtStsRpt = isomodels.FIToFIPmtStsRpt{
		GrpHdr:            isomodels.GrpHdrPacs002{},
		OrgnlGrpInfAndSts: mapOrgnlGrpInfAndSts(pos),
		TxInfAndSts:       mapTxInfAndSts(pos, status),
	}
	return &pacs002
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

func mapTxInfAndSts(pos []*models.PaymentOrder, status map[string]string) []*isomodels.TxInfAndSts {
	var txInfAndSts []*isomodels.TxInfAndSts
	stsId, _ := gonanoid.New(12)

	for _, po := range pos {
		if po != nil {
			settlmKey := sha256.Sum256([]byte(po.InstructionId + po.EndToEndId + po.TransactionId + po.UETR))
			txInfAndSts = append(txInfAndSts, &isomodels.TxInfAndSts{
				StsId:           stsId,
				OrgnlInstrId:    po.InstructionId,
				OrgnlEndToEndId: po.EndToEndId,
				OrgnlTxId:       po.TransactionId,
				OrgnlUETR:       po.UETR,
				TxSts:           status[po.TransactionId],
				OrgnlGrpInf: &isomodels.OrgnlGrpInf{
					OrgnlMsgId:   po.MessageId,
					OrgnlMsgNmId: po.MessageNameId,
					OrgnlCreDtTm: po.CreationDateTime,
				},
				ChrgsInf:    mapChrgsInf(po.Charges),
				AccptncDtTm: po.TxnAcceptanceDateTime,
				PrcgDt: &isomodels.DtAndDtTmChoice{
					DtTm: po.ProcessingDateTime,
				},
				FctvIntrBkSttlmDt: &isomodels.DtAndDtTmChoice{
					DtTm: po.EffectiveSettlementDateTime,
				},
				AcctSvcrRef: po.Id,
				ClrSysRef:   po.ClearingSystemReference,
				CdtSttlmKey: hex.EncodeToString(settlmKey[:]),
				OrgnlTxRef:  mapOrgnlTxRef(po),
				SplmtryData: mapSplmtryData(po.SupplementaryData),
			})
		}
	}
	return txInfAndSts
}

func mapSplmtryData(supplementaryData []*models.SupplementaryData) []*isomodels.SplmtryData {
	var supplData []*isomodels.SplmtryData
	for _, splData := range supplementaryData {
		if splData != nil {
			supplData = append(supplData, &isomodels.SplmtryData{
				PlcAndNm: splData.PlaceAndName,
				Envlp: &isomodels.Envlp{
					Data: splData.Envelope,
				},
			})
		}
	}
	return supplData
}

func mapChrgsInf(charges []*models.Charges) []*isomodels.ChrgsInf {
	var chrgsInf []*isomodels.ChrgsInf
	for _, chrgs := range charges {
		if chrgs != nil {
			chg := isomodels.ChrgsInf{
				Amt: mapPoAmount(chrgs.Amount),
				Agt: mapPoAgent(chrgs.Agent),
			}
			if chrgs.Type != nil {
				chg.Tp = &isomodels.ChrgsTp{}
				chg.Tp.Cd = chrgs.Type.Code
				if chrgs.Type.ProprietaryId != "" || chrgs.Type.ProprietaryIssuer != "" {
					chg.Tp.Prtry = &isomodels.ChrgsTpPrtry{}
					chg.Tp.Prtry.Id = chrgs.Type.ProprietaryId
					chg.Tp.Prtry.Issr = chrgs.Type.ProprietaryIssuer
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

	return &isomodels.Agent{
		FinInstnId: &isomodels.FinInstnId{
			BICFI: agent.FiiBicfi,
			ClrSysMmbId: &isomodels.ClrSysMmbId{
				ClrSysId: &isomodels.ClrSysId{
					Cd:    agent.FiiClearingSystemIdCode,
					Prtry: agent.FiiClearingSystemIdProprietary,
				},
				MmbId: agent.FiiMemberId,
			},
			LEI: agent.FiiLei,
			Nm:  agent.FiiName,
			Othr: &isomodels.FinInstnOthr{
				Id:   agent.FiiOtherId,
				Issr: agent.FiiOtherIssuer,
				SchmeNm: &isomodels.FinInstnSchmeNm{
					Cd:    agent.FiiOtherSchemeNameCode,
					Prtry: agent.FiiOtherSchemeNameProprietary,
				},
			},
		},
		BrnchId: &isomodels.BrnchId{
			Id:  agent.BiId,
			LEI: agent.BiLei,
			Nm:  agent.BiName,
		},
	}
}

func mapPstlAdr(postalAddress *models.PostalAddress) *isomodels.PstlAdr {
	return &isomodels.PstlAdr{
		AdrTp: &isomodels.AdrTp{
			Cd: postalAddress.AddressTypeCode,
			Prtry: &isomodels.AdrTpPrtry{
				Id:      postalAddress.AddressTypeProprietaryId,
				Issr:    postalAddress.AddressTypeProprietaryIssuer,
				SchmeNm: postalAddress.AddressTypeProprietarySchemeName,
			},
		},
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
}

func mapOrgnlTxRef(po *models.PaymentOrder) *isomodels.OrgnlTxRefPacs002 {
	if po == nil {
		return nil
	}

	orgnl := isomodels.OrgnlTxRefPacs002{
		IntrBkSttlmDt: po.SettlementDate,
		SttlmInf: &isomodels.SttlmInf{
			SttlmMtd:  po.SettlementMethod,
			SttlmAcct: mapPoAccount(po.SettlementAcct),
		},
		PmtTpInf: mapPmtTpInf(po.PaymentTypeInfo),
		PmtMtd:   "TRF",
		RmtInf: &isomodels.RmtInf{
			Ustrd: po.UnstructuredRemittanceInfo,
		},
		Dbtr: &isomodels.PartyOrAgentChoice{
			Pty: mapPoParty(po.Debtor),
		},
		UltmtDbtr: &isomodels.PartyOrAgentChoice{
			Pty: mapPoParty(po.UltimateDebtor),
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
		UltmtCdtr: &isomodels.PartyOrAgentChoice{
			Pty: mapPoParty(po.UltimateCreditor),
		},
		Purp: &isomodels.Purp{
			Cd:    po.PurposeCode,
			Prtry: po.PurposeProprietary,
		},
	}

	if po.SettlementAmount != nil {
		orgnl.IntrBkSttlmAmt = &isomodels.Amount{}
		orgnl.IntrBkSttlmAmt.Value = po.SettlementAmount.Value
		orgnl.IntrBkSttlmAmt.Currency = po.SettlementAmount.Currency
	}

	return &orgnl
}

func mapPoAccount(account *models.Account) *isomodels.Account {
	if account == nil {
		return nil
	}

	return &isomodels.Account{
		Id: &isomodels.AccountId{
			IBAN: account.Iban,
			Othr: &isomodels.AccountOthr{
				Id:   account.OtherId,
				Issr: account.OtherIssuer,
				SchmeNm: &isomodels.AccountSchmeNm{
					Cd:    account.OtherSchemeNameCode,
					Prtry: account.OtherSchemeNameProprietary,
				},
			},
		},
		Tp: &isomodels.AccountType{
			Cd:    account.TypeCode,
			Prtry: account.TypeProprietary,
		},
		Ccy: account.Currency,
		Nm:  account.Name,
		Prxy: &isomodels.AccountProxy{
			Tp: account.ProxyType,
			Id: account.ProxyId,
		},
	}
}

func mapPmtTpInf(pti *models.PaymentTypeInfo) *isomodels.PmtTpInfPacs002 {
	if pti == nil {
		return nil
	}

	return &isomodels.PmtTpInfPacs002{
		InstrPrty: pti.InstructionPriority,
		ClrChanl:  pti.ClearingChannel,
		SvcLvl:    mapServiceLevel(pti.ServiceLevel),
		LclInstrm: &isomodels.LclInstrm{
			Cd:    pti.LocalInstrumentCode,
			Prtry: pti.LocalInstrumentProprietary,
		},
		CtgyPurp: &isomodels.CtgyPurp{
			Cd:    pti.CategoryPurposeCode,
			Prtry: pti.CategoryPurposeProprietary,
		},
	}
}

func mapServiceLevel(serviceLevels []*models.CodeOrProprietary) []*isomodels.SvcLvl {
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

	return &isomodels.Party{
		Nm:      party.Name,
		PstlAdr: mapPstlAdr(party.PostalAddress),
		Id: &isomodels.PartyId{
			OrgId: &isomodels.OrgId{
				AnyBIC: party.OrgIdAnyBic,
				LEI:    party.OrgIdLei,
				Othr:   mapPoOrgIdOthr(party.OrgIdOther),
			},
			PrvtId: &isomodels.PrvtId{
				DtAndPlcOfBirth: &isomodels.DateAndPlaceOfBirth{
					BirthDt:     party.PrivateIdBirthDate,
					CityOfBirth: party.PrivateIdCityOfBirth,
					CtryOfBirth: party.PrivateIdCountryOfBirth,
					PrvcOfBirth: party.PrivateIdProvinceOfBirth,
				},
				Othr: mapPoPrvtIdOthr(party.PrivateIdOther),
			},
		},
		CtryOfRes: party.CountryOfResidence,
		CtctDtls:  mapCtctDtls(party.ContactDetails),
	}
}

func mapPoOrgIdOthr(orgIdOthers []*models.OrgIdOthr) []*isomodels.OrgIdOthr {
	var isomodelsOthers []*isomodels.OrgIdOthr
	for _, org := range orgIdOthers {
		isomodelsOthers = append(isomodelsOthers, &isomodels.OrgIdOthr{
			Id:   org.Id,
			Issr: org.Issuer,
			SchmeNm: &isomodels.OrgSchmeNm{
				Cd:    org.SchemeNameCode,
				Prtry: org.SchemeNameProprietary,
			},
		})
	}
	return isomodelsOthers
}

func mapPoPrvtIdOthr(privateIdOthers []*models.PrivateIdOthr) []*isomodels.PrvtIdOthr {
	var isoPrvtOthers []*isomodels.PrvtIdOthr
	for _, p := range privateIdOthers {
		isoPrvtOthers = append(isoPrvtOthers, &isomodels.PrvtIdOthr{
			Id:   p.Id,
			Issr: p.Issuer,
			SchmeNm: &isomodels.PersonSchmeNm{
				Cd:    p.SchemeNameCode,
				Prtry: p.SchemeNameProprietary,
			},
		})
	}
	return isoPrvtOthers
}

func mapCtctDtls(contact *models.ContactDetails) *isomodels.CtctDtls {
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
	var isoContacts []*isomodels.ContactOthr
	for _, c := range contactOthers {
		isoContacts = append(isoContacts, &isomodels.ContactOthr{
			Id:      c.Id,
			ChanlTp: c.ChannelType,
		})
	}
	return isoContacts
}
