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

# Import Libs required for Bearer Token OAuth2.0
from oauthlib.oauth2 import BackendApplicationClient
from requests.auth import HTTPBasicAuth
from requests_oauthlib import OAuth2Session

def GetToken(logger, client_id, secret_key, api_url="https://apis.druva.com"):
	"""
		param Logger logger
		param str client_id: part of API Credentials
		param str secret_ley: part of API Credentials
		param api_url
		Refer https://developer.druva.com/reference#reference-getting-started
	"""
	try:
		auth = HTTPBasicAuth(client_id, secret_key)
		client = BackendApplicationClient(client_id=client_id)
		oauth = OAuth2Session(client=client)
		response = oauth.fetch_token(token_url=api_url+'/token', auth=auth)
		auth_token = response['access_token']
		expires_at = response['expires_at']
	except Exception as e:
		logger.error("GetToken failed for url=%s, clientID=%s, Error=%s"%
	       (api_url+'/token', client_id, str(e)))
		raise e
	return auth_token, expires_at