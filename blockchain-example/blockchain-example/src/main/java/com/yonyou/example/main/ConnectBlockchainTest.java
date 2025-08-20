package com.yonyou.example.main;

import com.yonyou.example.baasnet.CustomClient;
import com.yonyou.example.baasnet.FabricTools;
import org.hyperledger.fabric.sdk.*;
import org.springframework.util.ResourceUtils;

import java.nio.file.Paths;
import java.util.Collection;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;
import java.util.concurrent.TimeUnit;

/**
 * @ClassName Test
 * @Description 调用智能合约测试
 * @Author 孙世江
 * @Data 2020/8/14 0014 上午 10:53
 * @Version 1.0
 **/
public class ConnectBlockchainTest {
    public static void main(String[] args) {
        try {
            String cfgPath = ResourceUtils.getURL("classpath:").getPath();
            String os = System.getProperty("os.name");
            if (os.toLowerCase().startsWith("win")) {
                cfgPath = ResourceUtils.getURL("classpath:").getPath().substring(1);
            }
            String netWorkPath = Paths.get(cfgPath, "network.json").toString();
            String certPath = Paths.get(cfgPath, "client.pem").toString();
            String privKeyPath = Paths.get(cfgPath, "private_key.pem").toString();

            CustomClient customClient = new CustomClient(netWorkPath, certPath, privKeyPath);
            Channel currentChannel = customClient.getChannel();
            HFClient hfClient = customClient.getHfClient();
            ChaincodeID ccId = customClient.buildChaincodeID();

            for (int i = 0; i < 3; i++) {
                int finalI = i;
                new Thread(() -> {
                    TransactionProposalRequest proposalRequest = hfClient.newTransactionProposalRequest();
                    proposalRequest.setChaincodeID(ccId);
                    proposalRequest.setFcn("set");
                    proposalRequest.setArgs("b"+finalI,"100"+finalI);
                    try {
                        Collection<ProposalResponse> proposalResponse = currentChannel.sendTransactionProposal(proposalRequest);
                        Collection<ProposalResponse> successful = new LinkedList<>();
                        Collection<ProposalResponse> failed = new LinkedList<>();
                        for (ProposalResponse response : proposalResponse) {
                            if (response.isVerified() && response.getStatus() == ProposalResponse.Status.SUCCESS) {
                                successful.add(response);
                            } else {
                                failed.add(response);
                            }
                        }
                        System.out.println("接收到 " + proposalResponse.size() + " 条提案结果. 成功并验证通过数量: " + successful.size() + " . 失败数量: " + failed.size());
                        if (failed.size() > 0) {
                            StringBuilder stringBuffer = new StringBuilder();
                            stringBuffer.append("error:");
                            for (ProposalResponse fail : failed) {
                                stringBuffer.append(fail.getMessage()).append(" The End! ");
                            }
                            System.out.println("错误信息 = " + stringBuffer.toString());
                        } else {
                            BlockEvent.TransactionEvent event = currentChannel.sendTransaction(proposalResponse).get(30, TimeUnit.SECONDS);
                            BlockEvent blockEvent = event.getBlockEvent();
                            List<Map<String, Object>> aaa =FabricTools.getRWSetFromBlock(blockEvent);
                            System.out.println("区块编号：" + event.getBlockEvent().getBlockNumber());
                            for (Map<String, Object> stringObjectMap : aaa) {
                                stringObjectMap.toString();
                            }
                        }
                    } catch (Exception e) {
                        e.printStackTrace();
                    }
                }).start();
            }
            Thread.sleep(3000);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
