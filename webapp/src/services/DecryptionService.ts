import * as cryptoJS from "crypto-js"

export class DecryptionService {
    private keyLength: number = 8;
    private numPBKDF2Iterations: number = 1024;

    private key: cryptoJS.lib.WordArray;
    private salt: cryptoJS.lib.WordArray;
    private iv: cryptoJS.lib.WordArray;
    private decryptor: any; // crypto.Cipher

    public constructor(password: string, salt: string, iv: string) {
        this.salt = cryptoJS.enc.Hex.parse(salt);
        this.iv = cryptoJS.enc.Hex.parse(iv);
        this.key = this.kdf(password, this.salt)

        this.createDecryptor() 
    }

    private kdf(password: string, salt: cryptoJS.lib.WordArray): cryptoJS.lib.WordArray {
        return cryptoJS.PBKDF2(password, salt, {
            keySize: this.keyLength, 
            iterations: this.numPBKDF2Iterations, 
            hasher: cryptoJS.algo.SHA256
        });
    }

    public createDecryptor() {
        this.decryptor = cryptoJS.algo.AES.createDecryptor(this.key, {
            mode:  cryptoJS.mode.CBC,
            iv: this.iv,
            padding: cryptoJS.pad.Pkcs7,
        });
    }

    public decrypt(encryptedStr: string): string {
        const encrypted = cryptoJS.enc.Hex.parse(encryptedStr);
        return this.decryptor.finalize(encrypted).toString(cryptoJS.enc.Latin1);
    }
}