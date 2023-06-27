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

import requests
import time
import json
import random

def get_retry_stats(logger, response, total_attempts, delay):
    if response.status_code == 200:
        return False, 0, 0
    elif response.status_code == 429:
        sleep_duration = delay + random.randint(1,30) + 1
        logger.info('Sleeping for %d seconds' %sleep_duration)
        time.sleep(sleep_duration)
        total_attempts = total_attempts - 1
        delay = 30 * (3-total_attempts)
        if total_attempts:
            return True, total_attempts, delay
    return False, 3, 0

def _get_api_call(logger, auth_token, api_url, api_path, headers, params, total_attempts, delay):
    """
        Invoke http GET request
    """
    logger.info('\nInvoking GET API call => %s' %(api_url+api_path))
    response = requests.get(api_url+api_path, headers=headers, params=params)

    do_retry, total_attempts, delay = get_retry_stats(logger, response, total_attempts, delay)
    if do_retry:
        response = _get_api_call(logger, auth_token, api_url, api_path, headers, params, total_attempts, delay)
        
    return response

def _post_api_call(logger, auth_token, api_url, api_path, headers, body, total_attempts, delay):
    """
        Invoke http POST request
    """
    logger.info('\nInvoking POST API call => %s' %(api_url+api_path))
    response = requests.post(api_url+api_path, headers=headers, data=json.dumps(body))

    do_retry, total_attempts, delay = get_retry_stats(logger, response, total_attempts, delay)
    if do_retry:
        response = _post_api_call(logger, auth_token, api_url, api_path, headers, body, total_attempts, delay)
        
    return response