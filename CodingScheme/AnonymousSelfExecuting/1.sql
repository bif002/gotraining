SELECT *
FROM (
    SELECT acct_no, To_char(txn_amt) AS TXN_AMT, msg_id, post_cd, post_dt, txn_cd, txn_typ_cd, txn_dr_cr_typ_cd, txn_delvry_chnl, en_txn_lng_desc_1, 
            en_txn_lng_desc_2, fr_txn_lng_desc_1, fr_txn_lng_desc_2, cba_txn_cd, ser_no
    FROM v_per_dda_txn
    WHERE acct_no = :1 AND post_dt BETWEEN :2 AND :3 ) A 
LEFT JOIN (
    SELECT acct_no AS ACC1, mrchnt_15, mrchnt_catg_cd, mrchnt_cntry_cd, 
        Row_number() OVER( partition BY mrchnt_15, mrchnt_catg_cd, mrchnt_cntry_cd ORDER BY '1') RN
    FROM v_per_dda_pos_txn
    WHERE acct_no = :4 AND txn_dt BETWEEN :5 AND :6 ) B 
ON A.en_txn_lng_desc_2 = B.mrchnt_15 AND A.acct_no = B.acc1 AND B.rn = 1


SELECT CUST_ID, ACCT_ID, ACCT_NO_MASKED, CLNT_PROD_CD, CURR_PROC_DT, OTPT_TXN_CD_INTRNL, TO_CHAR(TXN_AMT) AS TXN_AMT, TXN_POST_DT AS POST_DT, 
        TXN_DT, MRCHNT_DBA_CNTRY_CD, MRCHNT_DBA_NM, MRCHNT_DBA_CTY, MRCHNT_DBA_ST_CD, MRCHNT_SIC_CLSS_CD, TXN_CRNCY_CD, CHRGBK_IND, 
        POS_ENTRY_MODE_CD, DR_CR_IND, TXN_CATG_CD
FROM V_PER_5Q_CR_CRD
WHERE ACCT_NO_MASKED = :1 and ACCT_ID = :2 AND TXN_DT between :3 AND :4