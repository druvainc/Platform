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


try:
    #DCP: API call to get reports list.
    report_list_resp = reportingapi.GetReportsList(_logger, auth_token, api_url=api_url)
    if report_list_resp.status_code == 200:
        report_list = report_list_resp.json()
        print ('[RESPONSE]      :', json.dumps(report_list, indent=2))
    else:
        error_object = report_list_resp.json()
        raise Exception(error_object)
except Exception as e:
    _logger.error("Error in API call to get reports list => %s" %str(e))