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

from druvareportsdk.lib.restlib import restlib

reportingapi_base_path = "/platform/reporting/"
total_attempts = 3
delay = 0

def GetReportsList(logger, auth_token, product_ids_list="", version="v1", api_url="https://apis.druva.com"):
    """
        :param Logger logger
        :param str auth_token: access_token
        :param str product_ids_list: comma-separated list of product ids. Example- "8193,12289"
        :param str version
        :param str api_url
    """
    api_path = reportingapi_base_path + version + "/reports"

    headers = {'accept': 'application/json', 'Authorization': 'Bearer ' + auth_token}
    query_params = {'productIDs':product_ids_list}
    try:
        response = restlib._get_api_call(logger, auth_token, api_url, api_path, headers, query_params,
                                         total_attempts, delay)
    except Exception as e:
        logger.error("GetReportsList failed for url=%s, product_ids_list=%s, Error=%s" %
                     (api_url+api_path, product_ids_list, str(e)))
        raise e
    return response

def GetReportsData(logger, auth_token, report_id, body, version="v1", api_url="https://apis.druva.com"):
    """
        :param Logger logger
        :param str auth_token: access_token
        :param dict body: request body
        :param str version
        :param str api_url
    """
    api_path = reportingapi_base_path + version + "/reports/{report_id}".format(report_id=report_id)

    headers = {'accept': 'application/json', 'Authorization': 'Bearer ' + auth_token}
    try:
        response = restlib._post_api_call(logger, auth_token, api_url, api_path, headers, body,
                                          total_attempts, delay)
    except Exception as e:
        logger.error("GetReportsData failed for url=%s, request body=%r, Error=%s" %
                     (api_url+api_path, body, str(e)))
        raise e
    return response