#******************************************************************************
# Druva Confidential and Proprietary
#
#  Copyright (C) 2023-24, Druva Technologies Pte. Ltd.  ALL RIGHTS RESERVED.
#
#  Except as specifically permitted herein, no portion of the
#  information, including but not limited to object code and source
#  code, may be reproduced, modified, distributed, republished or
#  otherwise utilized in any form or by any means for any purpose
#  without the prior written permission of Druva Technologies Pte. Ltd.
#
#  Visit http://www.druva.com/ for more information.
#******************************************************************************

import json
import builtins
import logging

from druvareportsdk.authentication import authenticate
from druvareportsdk.reportingapi import reportingapi

builtins.globallogger = logging.getLogger('')
_logger = builtins.globallogger

# API credentials
client_id = ''
secret_key = ''

api_url = "https://apis-us0.druva.com"

# Fetch the token.
auth_token,_ = authenticate.GetToken(_logger, client_id, secret_key, api_url)
print('Auth_token: ', auth_token)

# The version supported for the report.
version = "v1"

# Unique reportID of the report.
report_id = "epLastBackupStatus"

body = {
    "filters":{
        "pageSize": 5,
        "filterBy": [{
            "fieldName": "status",
            "value": "Backup Failed",
            "operator": "EQUAL"
        }]
    }
}

reattempt = 1
while True:
    try:
        # DCP: API call to fetch reports data for failed backup operations.
        report_data_resp = reportingapi.GetReportsData(_logger, auth_token, report_id, body, version=version, api_url=api_url)
        if report_data_resp.status_code == 200:
            report_data = report_data_resp.json()
            print ('[RESPONSE]      :', json.dumps(report_data, indent=2))
            nextPageToken = report_data['nextPageToken']
            if not nextPageToken:
                break
            body['pageToken'] = nextPageToken
            reattempt = 1
        elif report_data_resp.status_code == 403 and reattempt:
            auth_token,_ = authenticate.GetToken(_logger, client_id, secret_key, api_url)
            reattempt = 0
        else:
            error_object = report_data_resp.json()
            raise Exception(error_object)
    except Exception as e:
        _logger.error("Error in API call to fetch reports data => %s" %str(e))
        break