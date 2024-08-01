package com.decryptimer.encrypt.model;

public class Algorithm {
    private String name;
    private SecretKey secretKey;

    public Algorithm(String name, SecretKey secretKey){
        this.name = name;
        this.secretKey = secretKey;
    }
    public String getName() {
        return name;
    }
    public void setName(String name) {
        this.name = name;
    }
    public SecretKey getSecretKey() {
        return secretKey;
    }
    public void setSecretKey(SecretKey secretKey) {
        this.secretKey = secretKey;
    }
}
