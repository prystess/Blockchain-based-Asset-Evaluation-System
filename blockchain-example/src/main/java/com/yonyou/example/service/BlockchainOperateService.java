package com.yonyou.example.service;

import com.yonyou.example.baasnet.CustomClient;
import org.apache.tomcat.util.buf.HexUtils;
import org.hyperledger.fabric.sdk.*;
import org.hyperledger.fabric.sdk.exception.InvalidArgumentException;
import org.hyperledger.fabric.sdk.exception.ProposalException;
import org.springframework.stereotype.Service;
import org.springframework.util.ResourceUtils;

import java.nio.file.Paths;
import java.util.Collection;
import java.util.HashMap;
import java.util.Map;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.TimeoutException;

/**
 * @ClassName BlockchainOperateService
 * @Description TODO
 * @Author 孙世江
 * @Data 2020/8/14 0014 上午 9:32
 * @Version 1.0
 **/
@Service
public class BlockchainOperateService {

    private static CustomClient customClient;
    private static Channel currentChannel;
    private static HFClient hfClient;
    private static ChaincodeID ccId;

    static {
        try {
            String cfgPath = ResourceUtils.getURL("classpath:").getPath();
            System.out.println(cfgPath);
            String os = System.getProperty("os.name");
            if (os.toLowerCase().startsWith("win")) {
                cfgPath = ResourceUtils.getURL("classpath:").getPath().substring(1);
            }
            String netWorkPath = Paths.get(cfgPath, "network (4).json").toString();
            String certPath = Paths.get(cfgPath, "client.ceshi002.yonyou.com.pem").toString();
            String privKeyPath = Paths.get(cfgPath, "client.ceshi002.yonyou.com.priv").toString();

            CustomClient customClient = new CustomClient(netWorkPath, certPath,privKeyPath);
            currentChannel = customClient.getChannel();
            hfClient = customClient.getHfClient();
            ccId = customClient.buildChaincodeID();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    public Object set(String operate, String... args) {
        TransactionProposalRequest proposalRequest = hfClient.newTransactionProposalRequest();
        proposalRequest.setChaincodeID(ccId);
        proposalRequest.setFcn(operate);
        proposalRequest.setArgs(args);
        BlockEvent.TransactionEvent event = null;
        try {
            Collection<ProposalResponse> proposalResponse = currentChannel.sendTransactionProposal(proposalRequest);
            event = currentChannel.sendTransaction(proposalResponse).get(30, TimeUnit.SECONDS);
            System.out.println("区块编号：" + event.getBlockEvent().getBlockNumber());
        } catch (ProposalException | InvalidArgumentException | InterruptedException | ExecutionException | TimeoutException e) {
            e.printStackTrace();
        }
        Map<String, Object> map = new HashMap<>(3);
        map.put("txid", event.getTransactionID());
        map.put("valid", event.isValid());
        map.put("datahash", HexUtils.toHexString(event.getBlockEvent().getBlock().getHeader().getDataHash().toByteArray()));
        return map;
    }

    public String get(String operate, String key) throws Exception {
        QueryByChaincodeRequest request = hfClient.newQueryProposalRequest();
        request.setChaincodeID(ccId);
        request.setFcn(operate);
        request.setArgs(key);
        ProposalResponse[] responseArray = currentChannel.queryByChaincode(request).toArray(new ProposalResponse[0]);
        String response = new String(responseArray[0].getChaincodeActionResponsePayload());
        System.out.println("查询结果： " + response);
        return response;
    }
}
