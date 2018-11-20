package model

const LOGIN_FAILED = 100

const REGISTER_FAILED_BLANK_FIELD_FOUND = 110
const REGISTER_FAILED_EMAIL_ALREADY_EXISTS = 111

const ORDER_FAILED_PRODUCT_NOT_AVAILABLE_FOR_RENTING = 120
const ORDER_FAILED_RENT_DURATION_EXCEEDS_PRODUCT_MAX_RENT_DURATION = 121
const ORDER_FAILED_PRODUCT_ID_NOT_FOUND = 122
const ORDER_FAILED_QUANTITY_EXCEEDS_PRODUCT_QUANTITY = 123
const ORDER_FAILED_BORROWER_ID_NOT_FOUND = 124
const ORDER_FAILED_RENT_DURATION_DOESNT_MEET_MINIMUM_RENT_DURATION = 125
const ORDER_FAILED_BORROWER_IS_THE_TENANT = 126

const PRODUCT_NOT_FOUND = 130