package com.decryptimer.encrypt.model;

public class SecretKey {
    private String content;
    private Long length;
    private Byte size;

    public SecretKey(String content, Long length, Byte size){
        this.content = content;
        this.length = length;
        this.size = size;
    }

    public String getContent() {
        return content;
    }

    public void setContent(String content) {
        this.content = content;
    }

    public Long getLength() {
        return length;
    }

    public void setLength(Long length) {
        this.length = length;
    }

    public Byte getSize() {
        return size;
    }

    public void setSize(Byte size) {
        this.size = size;
    }

}
