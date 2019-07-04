package main

var PROCESSING_LIMIT = 50

var DISB_BDTYPE = "BD"

var ASN_BDTYPE = "ASN"

var record_types_to_keys_map = []byte(`{

	"purchase_order": {
		
		"primary_key": ["record-type","poNumber"]
	},

	"invoice_payment": {
		
		"primary_key": ["record-type","id"]
	},
	
	"challan": {
		
		"primary_key": ["record-type","challanNumber"]
	},

	"factor_limit": {
		
		"primary_key": ["record-type","supplierName"]
	},

	"fund_requisition": {
		
		"primary_key": ["record-type","supplierId","id"]
	},

	"invoice": {
		
		"primary_key": ["record-type","invoiceNumber"]
	},

	"users": {
		
		"primary_key": ["record-type","username","contactNo","email"]
	},

	"organisation": {
		
		"primary_key": ["record-type","id"]
	},

	"workorder_payment":{
		"primary_key": ["record-type","id"]
	},

	"workorder_fund_requisition":{
		"primary_key": ["record-type","id"]
	},

	"workorder_limit":{
		"primary_key": ["record-type","supplierName"]
	}


}`)

// "additional_json":{"record-type":"purchase_order","status":"string","deliveryStatus":"string","paymentStatus":"string","manufacturerId":"number",
// 						"supplierId": "string","supplierAddrsId": "string","delivryAddrsId": "string","billingAddrsId": "string","product": "[]interface{}"},

var config_json_bytes = []byte(`{
  "invoke_purchase_order_insert": {
    "date_fields": [
      "createDate",
      "issueDate"
    ],
    "amount_fields": [
      "grndTotal",
      "advPaymnt",
      "netPayable",
      "settledAmt",
      "discountedAmt",
      "POPayableAmount",
      "POPaymntAmount"

    ],
    "mandatory_fields": [
      "id",
      "manufacturerName",
      "poNumber",
      "createDate",
      "issueDate",
      "supplierName",
      "supplierAddrs",
      "delivryAddrs",
      "billingAddrs"
    ],
    "operation": {
      "primitive": "invoke_insert_update"
    },
    "additional_json": {
      "record-type": "purchase_order",
      "docType": "purchase_order"
    },
    "default_fields": {
      "isActive": true
    },
    "fields_map": {
      "manufacturerId": "manufacturerId",
      "manufacturerName": "manufacturerName",
      "poNumber": "poNumber",
      "createDate": "createDate",
      "issueDate": "issueDate",
      "supplierId": "supplierId",
      "supplierName": "supplierName",
      "supplierAddrsId": "supplierAddrsId",
      "supplierAddrs": "supplierAddrs",
      "delivryAddrsId": "delivryAddrsId",
      "delivryAddrs": "delivryAddrs",
      "billingAddrsId": "billingAddrsId",
      "billingAddrs": "billingAddrs",
      "productDetails": [
        {
          "id": "id",
          "product": "product",
          "description": "description",
          "unitId": "unitId",
          "unitName": "unitName",
          "qty": "qty",
          "unitPrice": "unitPrice",
          "totalPrice": "totalPrice",
          "taxCodeId": "taxCodeId",
          "taxCodeName": "taxCodeName"
        }
      ],
      "grndTotal": "grndTotal",
      "advPaymnt": "advPaymnt",
      "netPayable": "netPayable",
      "amtInWord": "amtInWord",
      "status": "status",
      "docStatus": "docStatus",
      "reason": "reason",
      "deliveryStatus": "deliveryStatus",
      "paymentStatus": "paymentStatus",
      "finalChallan": "finalChallan",
      "version": "version",
      "id": "id",
      "createdBy": "createdBy",
      "createdOn": "createdOn",
      "checkedBy": "checkedBy",
      "checkedOn": "checkedOn",
      "authorisedBy": "authorisedBy",
      "authorisedOn": "authorisedOn",
      "rejectedBy": "rejectedBy",
      "rejectedOn": "rejectedOn",
      "apr": "apr",
      "loanLimit": "loanLimit",
      "reqAmount": "reqAmount",
      "discountedAmt": "discountedAmt",
      "POPayableAmount": "POPayableAmount",
      "POPaymntAmount": "POPaymntAmount",
      "settledAmt": "settledAmt"
    },
    "Validation_checks": {
      "PO_Validation": {
        "Target_addtional_primary_key_fields_values": {
          "record-type": "purchase_order"
        },
        "Target_primary_key": [
          "record-type",
          "poNumber"
        ]
      }
    }
  },
  "query_po_primary_key": {
    "operation": {
      "primitive": "query_primary_key"
    },
    "additional_json": {
      "record-type": "purchase_order"
    }
  },
  "query_po_by_manufacturer_id": {
    "operation": {
      "primitive": "query_using_rich_query"
    }
  },
  "query_po_by_supplier_id": {
    "operation": {
      "primitive": "query_using_rich_query"
    }
  },
  "query_po_primary_key_history": {
    "operation": {
      "primitive": "query_primary_key_history"
    },
    "additional_json": {
      "record-type": "purchase_order"
    }
  },
  "invoke_challan_insert": {
    "date_fields": [
      "challanDate",
      "createDate",
      "poDate"
    ],
    "mandatory_fields": [
      "supplier",
      "challanNumber",
      "challanDate",
      "manufacturer",
      "manufacturersAddress"
    ],
    "operation": {
      "primitive": "invoke_insert_update"
    },
    "additional_json": {
      "record-type": "challan",
      "docType": "challan"
    },
    "default_fields": {
      "isActive": true
    },
    "fields_map": {
      "supplierId": "supplierId",
      "supplier": "supplier",
      "challanDate": "challanDate",
      "challanNumber": "challanNumber",
      "manufacturer": "manufacturer",
      "manufacturerId": "manufacturerId",
      "manufacturersAddress": "manufacturersAddress",
      "poNumberId": "poNumberId",
      "poNumber": "poNumber",
      "poDate": "poDate",
      "referenceNo": "referenceNo",
      "product": [
        {
          "id": "id",
          "product": "product",
          "description": "description",
          "unitId": "unitId",
          "unitName": "unitName",
          "qty": "qty",
          "unitPrice": "unitPrice",
          "totalPrice": "totalPrice",
          "taxCodeId": "taxCodeId",
          "taxCodeName": "taxCodeName"
        }
      ],
      "preparedBy": "preparedBy",
      "designation": "designation",
      "deliveryDate": "deliveryDate",
      "deliveryAddress": "deliveryAddress",
      "finalChallan": "finalChallan",
      "status": "status",
      "reason": "reason",
      "createDate": "createDate",
      "version": "version",
      "id": "id",
      "createdBy": "createdBy",
      "createdOn": "createdOn",
      "checkedBy": "checkedBy",
      "checkedOn": "checkedOn",
      "authorisedBy": "authorisedBy",
      "authorisedOn": "authorisedOn",
      "rejectedBy": "rejectedBy",
      "rejectedOn": "rejectedOn"
    },
    "Validation_checks": {
      "Challan_Validation": {
        "Target_addtional_primary_key_fields_values": {
          "record-type": "challan"
        },
        "Target_primary_key": [
          "record-type",
          "challanNumber"
        ]
      }
    }
  },
  "query_challan_primary_key": {
    "operation": {
      "primitive": "query_primary_key"
    },
    "additional_json": {
      "record-type": "challan"
    }
  },
  "query_challan_primary_key_history": {
    "operation": {
      "primitive": "query_primary_key_history"
    },
    "additional_json": {
      "record-type": "challan"
    }
  },
  "query_challan_by_manufacturer_id": {
    "operation": {
      "primitive": "query_using_rich_query"
    }
  },
  "query_challan_by_supplier_id": {
    "operation": {
      "primitive": "query_using_rich_query"
    }
  },
  "invoke_factor_limit_insert": {
    "date_fields": [
      "effectiveDate",
      "expiryDate"
    ],
    "amount_fields": [
      "creditLimit",
      "totaladvPayment",
      "totalNetInvoice",
      "totalPayable",
      "SLLimit"
    ],
    "mandatory_fields": [
      "branchId",
      "loanAccNo",
      "supplierName",
      "supplierAddress",
      "creditLimit",
      "effectiveDate",
      "term",
      "termPeriodName",
      "expiryDate",
      "creditPeriod",
      "gracePeriod",
      "advPayment",
      "interestRate",
      "penaltyCharge",
      "serviceCharge"
    ],
    "operation": {
      "primitive": "invoke_insert_update"
    },
    "additional_json": {
      "record-type": "factor_limit",
      "docType": "factor_limit"
    },
    "default_fields": {
      "isActive": true
    },
    "fields_map": {
      "branchId": "branchId",
      "SLLimit": "SLLimit",
      "loanAccNo": "loanAccNo",
      "supplierId": "supplierId",
      "supplierName": "supplierName",
      "supplierAddress": "supplierAddress",
      "creditLimit": "creditLimit",
      "effectiveDate": "effectiveDate",
      "term": "term",
      "termPeriodId": "termPeriodId",
      "termPeriodName": "termPeriodName",
      "expiryDate": "expiryDate",
      "creditPeriod": "creditPeriod",
      "gracePeriod": "gracePeriod",
      "advPayment": "advPayment",
      "interestRate": "interestRate",
      "SNDRate": "SNDRate",
      "penaltyCharge": "penaltyCharge",
      "serviceCharge": "serviceCharge",
      "SNDThreshold": "SNDThreshold",
      "salesLimitSetUp": [
        {
          "manufacturerName": "manufacturerName",
          "salesLedgerLimit": "salesLedgerLimit",
          "salesCreditPeriod": "salesCreditPeriod",
          "salesGracePeriod": "salesGracePeriod",
          "advPaymentRate": "advPaymentRate",
          "permittedCreditLimit": "permittedCreditLimit",
          "maxDays": "maxDays",
          "id": "id",
          "version": "version"
        }
      ],
      "status": "status",
      "id": "id",
      "version": "version",
      "createdBy": "createdBy",
      "createdOn": "createdOn",
      "checkedBy": "checkedBy",
      "checkedOn": "checkedOn",
      "authorisedBy": "authorisedBy",
      "authorisedOn": "authorisedOn",
      "rejectedBy": "rejectedBy",
      "rejectedOn": "rejectedOn"
    },
    "Validation_checks": {
      "factor_limit_Validation": {
        "Target_addtional_primary_key_fields_values": {
          "record-type": "factor_limit"
        },
        "Target_primary_key": [
          "record-type",
          "supplierName"
        ]
      }
    }
  },
  "query_factor_limit_primary_key": {
    "operation": {
      "primitive": "query_primary_key"
    },
    "additional_json": {
      "record-type": "factor_limit"
    }
  },
  "query_factor_limit_created_by": {
    "operation": {
      "primitive": "query_using_rich_query"
    }
  },
  "query_factor_limit_primary_key_history": {
    "operation": {
      "primitive": "query_primary_key_history"
    },
    "additional_json": {
      "record-type": "factor_limit"
    }
  },
  "invoke_fund_requisition_insert": {
    "date_fields": [
      "reqDate"
    ],
    "amount_fields": [
      "loanBal",
      "totAvlbl",
      "requestedAmt",
      "totFinAmt"
    ],
    "mandatory_fields": [
      "supplierId",
      "status",
      "invoices"
    ],
    "operation": {
      "primitive": "invoke_insert_update"
    },
    "additional_json": {
      "record-type": "fund_requisition",
      "docType": "fund_requisition"
    },
    "fields_map": {
      "id": "id",
      "reqDate": "reqDate",
      "loanBal": "loanBal",
      "supplierId": "supplierId",
      "supplierName": "supplierName",
      "availableforFactoring": "availableforFactoring",
      "totAvlbl": "totAvlbl",
      "requestedAmt": "requestedAmt",
      "invoices": [
        {
          "grandTotal": "grandTotal",
          "totalVat": "totalVat",
          "totalTax": "totalTax",
          "netInvoiceAmount": "netInvoiceAmount",
          "advancePayment": "advancePayment",
          "totalPayable": "totalPayable",
          "invoiceTax": "invoiceTax",
          "invoiceVat": "invoiceVat",
          "invoiceDeduction": "invoiceDeduction",
          "invoicePayableAmount": "invoicePayableAmount",
          "unadjustedAmount": "unadjustedAmount",
          "advanceAmt": "advanceAmt",
          "financeAmt": "financeAmt",
          "financedAmt": "financedAmt",
          "availableLimit": "availableLimit",
          "slAvailable": "slAvailable",
          "reqAmount": "reqAmount",
          "apr": "apr",
          "loanLimit": "loanLimit",
          "discountAmt": "discountAmt",
          "supplierName": "supplierName",
          "supplierId": "supplierId",
          "invoiceDate": "invoiceDate",
          "invoiceNumber": "invoiceNumber",
          "manufacturer": "manufacturer",
          "manufacturerId": "manufacturerId",
          "manufacturersAddress": "manufacturersAddress",
          "poNumberId": "poNumberId",
          "poNumber": "poNumber",
          "poDate": "poDate",
          "challanId": "challanId",
          "challanNumber": "challanNumber",
          "challanDate": "challanDate",
          "referenceNo": "referenceNo",
          "productDetails": [
            {
              "id": "id",
              "product": "product",
              "quantity": "quantity",
              "unitId": "unitId",
              "unitName": "unitName",
              "unitPrice": "unitPrice",
              "totalPrice": "totalPrice",
              "vat": "vat",
              "tax": "tax",
              "netAmount": "netAmount"
            }
          ],
          "amtInWord": "amtInWord",
          "preparedBy": "preparedBy",
          "designation": "designation",
          "status": "status",
          "paymentStatus": "paymentStatus",
          "reason": "reason",
          "date": "date",
          "createDate": "createDate",
          "expiredDate": "expiredDate",
          "pendingAcceptance": "pendingAcceptance",
          "discountedAmt": "discountedAmt",
          "settledAmt": "settledAmt",
          "version": "version",
          "id": "id"
        }
      ],
      "status": "status",
      "totFinAmt": "totFinAmt",
      "pendingAcceptance": "pendingAcceptance",
      "version": "version",
      "createdBy": "createdBy",
      "createdOn": "createdOn",
      "checkedBy": "checkedBy",
      "checkedOn": "checkedOn",
      "authorisedBy": "authorisedBy",
      "authorisedOn": "authorisedOn",
      "rejectedBy": "rejectedBy",
      "rejectedOn": "rejectedOn"
    }
  },
  "query_fund_requisition_primary_key": {
    "operation": {
      "primitive": "query_primary_key"
    },
    "additional_json": {
      "record-type": "fund_requisition"
    }
  },
  "query_fund_requisition_by_supplier_id": {
    "operation": {
      "primitive": "query_using_rich_query"
    }
  },
  "query_fund_requisition_primary_key_history": {
    "operation": {
      "primitive": "query_primary_key_history"
    },
    "additional_json": {
      "record-type": "fund_requisition"
    }
  },
  "invoke_invoice_insert": {
    "amount_fields": [
      "grandTotal",
      "netInvoiceAmount",
      "advancePayment",
      "totalPayable",
      "invoiceDeduction",
      "invoicePayableAmount",
      "unadjustedAmount",
      "advanceAmt",
      "financeAmt",
      "financedAmt",
      "availableLimit",
      "slAvailable",
      "reqAmount",
      "loanLimit",
      "discountAmt",
      "discountedAmt",
      "pendingAcceptance",
      "settledAmt"
    ],
    "date_fields": [
      "invoiceDate",
      "poDate",
      "challanDate"
    ],
    "mandatory_fields": [
      "supplierName",
      "invoiceDate",
      "invoiceNumber",
      "manufacturer",
      "manufacturersAddress"
    ],
    "operation": {
      "primitive": "invoke_insert_update"
    },
    "additional_json": {
      "record-type": "invoice",
      "docType": "invoice"
    },
    "default_fields": {
      "isActive": true
    },
    "fields_map": {
      "grandTotal": "grandTotal",
      "totalVat": "totalVat",
      "totalTax": "totalTax",
      "netInvoiceAmount": "netInvoiceAmount",
      "advancePayment": "advancePayment",
      "totalPayable": "totalPayable",
      "invoiceTax": "invoiceTax",
      "invoiceVat": "invoiceVat",
      "invoiceDeduction": "invoiceDeduction",
      "invoicePayableAmount": "invoicePayableAmount",
      "unadjustedAmount": "unadjustedAmount",
      "advanceAmt": "advanceAmt",
      "financeAmt": "financeAmt",
      "financedAmt": "financedAmt",
      "availableLimit": "availableLimit",
      "slAvailable": "slAvailable",
      "reqAmount": "reqAmount",
      "apr": "apr",
      "loanLimit": "loanLimit",
      "discountAmt": "discountAmt",
      "supplierName": "supplierName",
      "supplierId": "supplierId",
      "invoiceDate": "invoiceDate",
      "invoiceNumber": "invoiceNumber",
      "manufacturer": "manufacturer",
      "manufacturerId": "manufacturerId",
      "manufacturersAddress": "manufacturersAddress",
      "poNumberId": "poNumberId",
      "poNumber": "poNumber",
      "poDate": "poDate",
      "challanId": "challanId",
      "challanNumber": "challanNumber",
      "challanDate": "challanDate",
      "referenceNo": "referenceNo",
      "productDetails": [
        {
          "id": "id",
          "product": "product",
          "quantity": "quantity",
          "unitId": "unitId",
          "unitName": "unitName",
          "unitPrice": "unitPrice",
          "totalPrice": "totalPrice",
          "vat": "vat",
          "tax": "tax",
          "netAmount": "netAmount"
        }
      ],
      "amtInWord": "amtInWord",
      "preparedBy": "preparedBy",
      "designation": "designation",
      "status": "status",
      "paymentStatus": "paymentStatus",
      "reason": "reason",
      "date": "date",
      "createDate": "createDate",
      "expiredDate": "expiredDate",
      "pendingAcceptance": "pendingAcceptance",
      "discountedAmt": "discountedAmt",
      "settledAmt": "settledAmt",
      "version": "version",
      "id": "id",
      "createdBy": "createdBy",
      "createdOn": "createdOn",
      "checkedBy": "checkedBy",
      "checkedOn": "checkedOn",
      "authorisedBy": "authorisedBy",
      "authorisedOn": "authorisedOn",
      "rejectedBy": "rejectedBy",
      "rejectedOn": "rejectedOn"
    },
    "Validation_checks": {
      "Invoice_Validation": {
        "Target_addtional_primary_key_fields_values": {
          "record-type": "invoice"
        },
        "Target_primary_key": [
          "record-type",
          "invoiceNumber"
        ]
      }
    }
  },
  "query_invoice_primary_key": {
    "operation": {
      "primitive": "query_primary_key"
    },
    "additional_json": {
      "record-type": "invoice"
    }
  },
  "query_invoice_by_supplier_id": {
    "operation": {
      "primitive": "query_using_rich_query"
    }
  },
  "query_invoice_by_manufacturer_id": {
    "operation": {
      "primitive": "query_using_rich_query"
    }
  },
  "query_invoice_primary_key_history": {
    "operation": {
      "primitive": "query_primary_key_history"
    },
    "additional_json": {
      "record-type": "invoice"
    }
  },
  "invoke_users_insert": {
    "mandatory_fields": [
      "firstName",
      "lastName",
      "id",
      "designation",
      "roleName",
      "userTypeName",
      "contactNo",
      "email",
      "orgId",
      "newUser"
    ],
    "operation": {
      "primitive": "invoke_insert_update"
    },
    "additional_json": {
      "record-type": "users",
      "docType": "users"
    },
    "fields_map": {
      "id": "id",
      "orgName": "orgName",
      "orgId": "orgId",
      "firstName": "firstName",
      "lastName": "lastName",
      "designation": "designation",
      "status": "status",
      "roleId": "roleId",
      "roleName": "roleName",
      "contactNo": "contactNo",
      "email": "email",
      "userTypeId": "userTypeId",
      "userTypeName": "userTypeName",
      "username": "username",
      "hash": "hash",
      "reason": "reason",
      "version": "version",
      "createdBy": "createdBy",
      "createdOn": "createdOn",
      "checkedBy": "checkedBy",
      "checkedOn": "checkedOn",
      "authorisedBy": "authorisedBy",
      "authorisedOn": "authorisedOn",
      "rejectedBy": "rejectedBy",
      "rejectedOn": "rejectedOn",
      "newUser":"newUser"
    },
    "Validation_checks": {
      "Users_Validation": {
        "Target_addtional_primary_key_fields_values": {
          "record-type": "users"
        },
        "Target_primary_key": [
          "record-type",
          "username",
          "contactNo",
          "email"
        ]
      }
    }
  },
  "query_users_primary_key": {
    "operation": {
      "primitive": "query_primary_key"
    },
    "additional_json": {
      "record-type": "users"
    }
  },
  "query_users_by_org_id": {
    "operation": {
      "primitive": "query_using_rich_query"
    }
  },
  "query_users_primary_key_history": {
    "operation": {
      "primitive": "query_primary_key_history"
    },
    "additional_json": {
      "record-type": "users"
    }
  },
  "invoke_organisation_insert": {
    "date_fields": [],
    "mandatory_fields": [
      "designation",
      "mobile",
      "orgName",
      "phone",
      "email"
    ],
    "operation": {
      "primitive": "invoke_insert_update"
    },
    "additional_json": {
      "record-type": "organisation",
      "docType": "organisation"
    },
    "default_fields": {
      "isActive": true
    },
    "fields_map": {
      "id": "id",
      "entityTypeId": "entityTypeId",
      "entityTypeName": "entityTypeName",
      "orgName": "orgName",
      "natureOfBusinessId": "natureOfBusinessId",
      "natureOfBusinessName": "natureOfBusinessName",
      "contactName": "contactName",
      "designation": "designation",
      "mobile": "mobile",
      "phone": "phone",
      "email": "email",
      "isBank": "isBank",
      "address": [
        {
          "id": "id",
          "addressTypeId": "addressTypeId",
          "addressTypeName": "addressTypeName",
          "addressLine1": "addressLine1",
          "addressLine2": "addressLine2",
          "divisionId": "divisionId",
          "divisionName": "divisionName",
          "districtId": "districtId",
          "districtName": "districtName",
          "thanaId": "thanaId",
          "thanaName": "thanaName",
          "isMailingAddress": "isMailingAddress"
        }
      ],
      "company": [
        {
          "id": "id",
          "registrationTypeId": "registrationTypeId",
          "registrationTypeName": "registrationTypeName",
          "regNo": "regNo",
          "regDate": "regDate",
          "issuingCountryId": "issuingCountryId",
          "issuingCountryName": "issuingCountryName",
          "issuingOfficeId": "issuingOfficeId",
          "issuingOfficeName": "issuingOfficeName",
          "hasValidity": "hasValidity"
        }
      ],
      "shareholder": [
        {
          "id": "id",
          "shName": "shName",
          "roleId": "roleId",
          "roleName": "roleName",
          "perShare": "perShare",
          "idDocTypeId": "idDocTypeId",
          "idDocTypeName": "idDocTypeName",
          "shDocumentNo": "shDocumentNo",
          "issuingCountryId": "issuingCountryId",
          "issuingCountryName": "issuingCountryName",
          "shHasValidity": "shHasValidity"
        }
      ],
      "banks": [
        {
          "id": "id",
          "accountTitle": "accountTitle",
          "accountNumber": "accountNumber",
          "bankName": "bankName",
          "bankNameId": "bankNameId",
          "bankBranchName": "bankBranchName",
          "bankBranchId": "bankBranchId"
        }
      ],
      "status": "status",
      "reason": "reason",
      "createdBy": "createdBy",
      "createdOn": "createdOn",
      "checkedBy": "checkedBy",
      "checkedOn": "checkedOn",
      "authorisedBy": "authorisedBy",
      "authorisedOn": "authorisedOn",
      "rejectedBy": "rejectedBy",
      "rejectedOn": "rejectedOn"
    },
    "Validation_checks": {
      "Organisation_Validation": {
        "Target_addtional_primary_key_fields_values": {
          "record-type": "organisation"
        },
        "Target_primary_key": [
          "record-type",
          "id"
        ]
      }
    }
  },
  "query_organisation_primary_key": {
    "operation": {
      "primitive": "query_primary_key"
    },
    "additional_json": {
      "record-type": "organisation"
    }
  },
  "query_all": {
    "operation": {
      "primitive": "query_all_rich_query"
    }
  },
  "query_by_id_and_status": {
    "operation": {
      "primitive": "query_by_id_and_status"
    }
  },
  "query_by_user": {
    "operation": {
      "primitive": "query_by_user"
    }
  },
  "query_organisation_primary_key_history": {
    "operation": {
      "primitive": "query_primary_key_history"
    },
    "additional_json": {
      "record-type": "organisation"
    }
  },
  "invoke_invoice_payment_insert": {
    "date_fields": [
      "postingDate"
    ],
    "amount_fields": [
      "totalInvoiceAmount",
      "totaldeduction",
      "totaladvPayment",
      "totalNetInvoice",
      "totalPayable",
      "totalSettledAmt",
      "totalPaymntAmount"
    ],
    "mandatory_fields": [],
    "operation": {
      "primitive": "invoke_insert_update"
    },
    "additional_json": {
      "record-type": "invoice_payment",
      "docType": "invoice_payment"
    },
    "fields_map": {
      "invoicePayableDetails": [
        {
          "grandTotal": "grandTotal",
          "totalVat": "totalVat",
          "totalTax": "totalTax",
          "netInvoiceAmount": "netInvoiceAmount",
          "advancePayment": "advancePayment",
          "totalPayable": "totalPayable",
          "invoiceTax": "invoiceTax",
          "invoiceVat": "invoiceVat",
          "invoiceDeduction": "invoiceDeduction",
          "invoicePayableAmount": "invoicePayableAmount",
          "invoicePaidAmount": "invoicePaidAmount",
          "unadjustedAmount": "unadjustedAmount",
          "advanceAmt": "advanceAmt",
          "financeAmt": "financeAmt",
          "financedAmt": "financedAmt",
          "availableLimit": "availableLimit",
          "slAvailable": "slAvailable",
          "reqAmount": "reqAmount",
          "apr": "apr",
          "loanLimit": "loanLimit",
          "discountedAmt": "discountedAmt",
          "id": "id",
          "supplierName": "supplierName",
          "supplierId": "supplierId",
          "invoiceDate": "invoiceDate",
          "invoiceNumber": "invoiceNumber",
          "manufacturer": "manufacturer",
          "manufacturerId": "manufacturerId",
          "manufacturersAddress": "manufacturersAddress",
          "poNumberId": "poNumberId",
          "poNumber": "poNumber",
          "poDate": "poDate",
          "challanId": "challanId",
          "challanNumber": "challanNumber",
          "challanDate": "challanDate",
          "referenceNo": "referenceNo",
          "productDetails": [
            {
              "id": "string",
              "product": "string",
              "quantity": 0,
              "unitId": 0,
              "unitName": "string",
              "unitPrice": 0,
              "totalPrice": 0,
              "vat": 0,
              "tax": 0,
              "version": "string"
            }
          ],
          "amtInWord": "amtInWord",
          "preparedBy": "preparedBy",
          "designation": "designation",
          "status": "status",
          "paymentStatus": "paymentStatus",
          "reason": "reason"
        }
      ],
      "supplierName": "supplierName",
      "supplierId": "supplierId",
      "manufacturerId": "manufacturerId",
      "poNumber": "poNumber",
      "poNumberId": "poNumberId",
      "totalInvoiceAmount": "totalInvoiceAmount",
      "totaladvPayment": "totaladvPayment",
      "totalNetInvoice": "totalNetInvoice",
      "totalTax": "totalTax",
      "totalVat": "totalVat",
      "totaldeduction": "totaldeduction",
      "totalPaidAmount": "totalPaidAmount",
      "totalPayable": "totalPayable",
      "fiBankAccNo": "fiBankAccNo",
      "fiBankName": "fiBankName",
      "fiBankDtlsId": "fiBankDtlsId",
      "fiBankId": "fiBankId",
      "fiId": "fiId",
      "fiName": "fiName",
      "manfBankAccNo": "manfBankAccNo",
      "manfBankName": "manfBankName",
      "manfBankNameId": "manfBankNameId",
      "manfBankId": "manfBankId",
      "postingDate": "postingDate",
      "cbsbatchno": "cbsbatchno",
      "cbstrackerno": "cbstrackerno",
      "status": "status",
      "reason": "reason",
      "cbsStatus": "cbsStatus",
      "version": "version",
      "id": "id",
      "totalPaymntAmount": "totalPaymntAmount",
      "totalSettledAmt": "totalSettledAmt"
    },
    "Validation_checks": {
      "invoice_payment_Validation": {
        "Target_addtional_primary_key_fields_values": {
          "record-type": "invoice_payment"
        },
        "Target_primary_key": [
          "record-type",
          "id"
        ]
      }
    }
  },
  "invoke_workorder_payment_insert": {
    "date_fields": [
      "postingDate"
    ],
    "amount_fields": [
      "totalWorkOrderAmount",
      "totaladvPayment",
      "netPayable",
      "totalSettledAmt",
      "totaldeduction",
      "totalPaymntAmount",
      "totalNetPayable"
    ],
    "mandatory_fields": [
      "id"
    ],
    "operation": {
      "primitive": "invoke_insert_update"
    },
    "additional_json": {
      "record-type": "workorder_payment","docType": "workorder_payment"
    },
    "fields_map": {
      "id": "id",
      "supplierName": "supplierName",
      "supplierId": "supplierId",
      "manufacturerId": "manufacturerId",
      "poNumber": "poNumber",
      "poNumberId": "poNumberId",
      "workOrderPayableDetails": [
        "string"
      ],
      "totalWorkOrderAmount": "totalWorkOrderAmount",
      "totaladvPayment": "totaladvPayment",
      "netPayable": "netPayable",
      "totalSettledAmt": "totalSettledAmt",
      "totalTax": "totalTax",
      "totalVat": "totalVat",
      "totaldeduction": "totaldeduction",
      "totalPaymntAmount": "totalPaymntAmount",
      "totalNetPayable": "totalNetPayable",
      "fiBankAccNo": "fiBankAccNo",
      "fiBankName": "fiBankName",
      "fiBankId": "fiBankId",
      "fiBankDtlsId": "fiBankDtlsId",
      "fiId": "fiId",
      "fiName": "fiName",
      "manfBankAccNo": "manfBankAccNo",
      "manfBankName": "manfBankName",
      "manfBankNameId": "manfBankNameId",
      "manfBankId": "manfBankId",
      "postingDate": "postingDate",
      "cbsbatchno": "cbsbatchno",
      "cbstrackerno": "cbstrackerno",
      "status": "status",
      "reason": "reason",
      "cbsStatus": "cbsStatus",
      "createdBy": "createdBy",
      "createdOn": "createdOn",
      "checkedBy": "checkedBy",
      "checkedOn": "checkedOn",
      "authorisedBy": "authorisedBy",
      "authorisedOn": "authorisedOn",
      "rejectedBy": "rejectedBy",
      "rejectedOn": "rejectedOn"
    },
    "Validation_checks": {
      "workorder_payment_Validation": {
        "Target_addtional_primary_key_fields_values": {
          "record-type": "workorder_payment"
        },
        "Target_primary_key": [
          "record-type",
          "id"
        ]
      }
    }
  },
  "invoke_workorder_fund_requisition_insert": {
    "date_fields": [
      "reqDate"
    ],
    "amount_fields": [
      "avlFund",
      "pendingAcceptance",
      "requestedAmt",
      "totDiscAmt",
      "totFactReq"
    ],
    "mandatory_fields": [
      "id"
    ],
    "operation": {
      "primitive": "invoke_insert_update"
    },
    "additional_json": {
      "record-type": "workorder_fund_requisition","docType": "workorder_fund_requisition"
    },
    "fields_map": {
      "id": "id",
      "avlFund": "avlFund",
      "pendingAcceptance": "pendingAcceptance",
      "requestedAmt": "requestedAmt",
      "totDiscAmt": "totDiscAmt",
      "supplierId": "supplierId",
      "supplierName": "supplierName",
      "productDetails": [
        "string"
      ],
      "status": "status",
      "reqDate": "reqDate",
      "totFactReq": "totFactReq",
      "reason": "reason",
      "createdBy": "createdBy",
      "createdOn": "createdOn",
      "checkedBy": "checkedBy",
      "checkedOn": "checkedOn",
      "authorisedBy": "authorisedBy",
      "authorisedOn": "authorisedOn",
      "rejectedBy": "rejectedBy",
      "rejectedOn": "rejectedOn"
    },
    "Validation_checks": {
      "workorder_fund_requisition_Validation": {
        "Target_addtional_primary_key_fields_values": {
          "record-type": "workorder_fund_requisition"
        },
        "Target_primary_key": [
          "record-type",
          "id"
        ]
      }
    }
  },
  "invoke_workorder_limit_insert": {
    "date_fields": [
      "effectiveDate",
      "expiryDate"
    ],
    "amount_fields": [
      "creditLimit",
      "advPayment"
    ],
    "mandatory_fields": [
      "branchId",
      "loanAccNo",
      "supplierName",
      "supplierAddress",
      "creditLimit",
      "effectiveDate",
      "term",
      "termPeriodName",
      "expiryDate",
      "creditPeriod",
      "gracePeriod",
      "advPayment",
      "interestRate",
      "penaltyCharge",
      "serviceCharge"
    ],
    "operation": {
      "primitive": "invoke_insert_update"
    },
    "additional_json": {
      "record-type": "workorder_limit","docType": "workorder_limit"
    },
    "default_fields": {
      "isActive": "True"
    },
    "fields_map": {
      "supplierAddress": "supplierAddress",
      "supplierId": "supplierId",
      "creditLimit": "creditLimit",
      "effectiveDate": "effectiveDate",
      "term": "term",
      "supplierName": "supplierName",
      "expiryDate": "expiryDate",
      "advPayment": "advPayment",
      "interestRate": "interestRate",
      "penaltyCharge": "penaltyCharge",
      "serviceCharge": "serviceCharge",
      "creditPeriod": "creditPeriod",
      "gracePeriod": "gracePeriod",
      "branchId": "branchId",
      "loanAccNo": "loanAccNo",
      "status": "status",
      "id": "id",
      "SLLimit": "SLLimit",
      "termPeriodId": "termPeriodId",
      "termPeriodName": "termPeriodName",
      "debtorLimitSetUp": [
        "string"
      ],
      "version": "version",
      "createdBy": "createdBy",
      "createdOn": "createdOn",
      "checkedBy": "checkedBy",
      "checkedOn": "checkedOn",
      "authorisedBy": "authorisedBy",
      "authorisedOn": "authorisedOn",
      "rejectedBy": "rejectedBy",
      "rejectedOn": "rejectedOn"
    },
    "Validation_checks": {
      "workorder_limit_Validation": {
        "Target_addtional_primary_key_fields_values": {
          "record-type": "workorder_limit"
        },
        "Target_primary_key": [
          "record-type",
          "supplierName"
        ]
      }
    }
  }
}`)
