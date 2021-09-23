delpaths([
    ["paths", "/v1/charge/qrcode"],
    ["definitions", "newbillingGetChargeQrcodeResponse"],

    ["paths", "/v1/charge_returns/alipay"],
    ["definitions", "newbillingChargeAlipayReturnRequest"],
    ["definitions", "newbillingChargeAlipayReturnResponse"],

    ["paths", "/v1/recharges"],
    ["definitions", "newbillingCreateRechargeRequest"],
    ["definitions", "newbillingCreateRechargeResponse"],

    ["paths", "/v1/openapi/token:refresh"],
    ["definitions", "newbillingRefreshOpenApiTokenRequest"],
    ["definitions", "newbillingRefreshOpenApiTokenResponse"],

    ["paths", "/v1/openapi/token:refresh"],
    ["definitions", "newbillingRefreshOpenApiTokenRequest"],
    ["definitions", "newbillingRefreshOpenApiTokenRequest"],

    ["paths", "/v1/open/token:refresh"],
    ["definitions", "newbillingRefreshOpenTokenRequest"],
    ["definitions", "newbillingRefreshOpenTokenResponse"],

    ["paths", "/v1/open/token:delete"],
    ["definitions", "newbillingDeleteOpenTokenResponse"],
    ["definitions", "newbillingDeleteOpenTokenRequest"],

    ["paths", "/v1/email"],
    ["definitions", "newbillingSendMailsRequest"],
    ["definitions", "newbillingSendMailsResponse"],

    ["paths", "/v1/emailserviceconfig"],
    ["definitions", "newbillingGetEmailServiceConfigResponse"],
    ["definitions", "newbillingSetEmailServiceConfigResponse"],
    ["definitions", "newbillingSetEmailServiceConfigRequest"],


    ["paths", "/v1/messages"],
    ["definitions", "newbillingSendMessageRequest"],
    ["definitions", "newbillingSendMessageResponse"],

    ["paths", "/v1/notify/alipay"],
    ["definitions", "newbillingNotifyAlipayRequest"],
    ["definitions", "newbillingNotifyAlipayResponse"],

    ["paths", "/v1/notify/stripe/{access_sys_id}"],
    ["definitions", "newbillingNotifyStripeRequest"],
    ["definitions", "newbillingNotifyStripeResponse"],

    ["paths", "/v1/open/accesssystemcatalogs"],
    ["definitions", "newbillingCreateAccessSystemForOpenRequest"],
    ["definitions", "newbillingCreateAccessSystemForOpenResponse"],

    ["paths", "/v1/cost"],
    ["definitions", "newbillingCalculateProductPriceRequest"],
    ["definitions", "newbillingCalculateProductPriceResponse"],

    ["paths", "/v1/accesssystemcatalogs"],
    ["definitions", "newbillingCalculateProductPriceRequest"],
    ["definitions", "newbillingCreateAccessSystemResponse"],

    ["paths", "/v1/accesssystems"],
    ["definitions", "newbillingDescribeAccessSystemsResponse"],
    ["definitions", "newbillingDeleteAccessSystemsResponse"],
    ["definitions", "newbillingDeleteAccessSystemsRequest"],
    ["definitions", "newbillingModifyAccessSystemRequest"],
    ["definitions", "newbillingModifyAccessSystemResponse"],

    ["paths", "/v1/billingjobs"],
    ["definitions", "newbillingDescribeBillingJobsResponse"],
    ["definitions", "newbillingPerformBillingJobResponse"],
    ["definitions", "newbillingPerformBillingJoRequest"],

    ["paths", "/v1/packageorders", "post"],
    ["definitions", "newbillingCreatePrdOrderRequest"],

    ["paths", "/v1/charges", "post"],
    ["definitions", "newbillingCreateChargeRequest"],
    ["definitions", "newbillingCreateChargeResponse"],

    ["paths", "/v1/actions", "post"],
    ["definitions", "newbillingCreateActionResponse"],
    ["definitions", "newbillingCreateActionRequest"],

    ["paths", "/v1/actions", "patch"],
    ["definitions", "newbillingModifyActionResponse"],
    ["definitions", "newbillingModifyActionRequest"],

    ["paths", "/v1/actions", "delete"],
    ["definitions", "newbillingDeleteActionResponse"],
    ["definitions", "newbillingDeleteActionRequest"],

    ["paths", "/v1/collect/events", "post"],
    ["definitions", "newbillingCreateCollectEventResponse"],
    ["definitions", "newbillingCreateCollectEventRequest"],

    ["paths", "/v1/collect:start"],
    ["definitions", "newbillingStartCollectResponse"],
    ["definitions", "newbillingStartCollectRequest"],

    ["paths", "/v1/collect:stop"],
    ["definitions", "newbillingStopCollectResponse"],
    ["definitions", "newbillingStopCollectRequest"],

    ["paths", "/v1/collect/data"],
    ["definitions", "newbillingCreateCollectDataResponse"],
    ["definitions", "newbillingCreateCollectDataRequest"],
    ["definitions", "newbillingDescribeCollectDataResponse"],
    ["definitions", "newbillingDescribeCollectDataRequest"]
 ])