package com.decryptimer.encrypt.model;

public class Message {
    private String content;
    private String algorithm;
    private Boolean isEncrypted;
    private Byte size;

    public Message(String content, String algorithm, Boolean isEncrypted, Byte size){
        this.content = content;
        this.algorithm = algorithm;
        this.isEncrypted= isEncrypted;
        this.size = size;
    }

    public void setContent(String content){
        this.content = content;
    }
    public void setAlgorithm(String algorithm){
        this.algorithm = algorithm;
    }
    public void setIsEncrypted(Boolean isEncrypted){
        this.isEncrypted = isEncrypted;
    }
    public String getContent(){
        return this.content;
    }
    public String getAlgorithm(){
        return this.algorithm;
    }
    public Boolean getIsEncrypted(){
        return this.isEncrypted;
    }
    public Byte getSize() {
        return size;
    }

    public void setSize(Byte size) {
        this.size = size;
    }
}
