import * as cryptoJS from "crypto-js"

export class DecryptionService {
    keyLength: number = 8;
    numPBKDF2Iterations: number = 1024;
    decryptor: any; // crypto.Cipher

    public constructor(password: string, salt: string, iv: string) {
        const key = this.kdf(password, cryptoJS.enc.Hex.parse(salt))

        this.decryptor = cryptoJS.algo.AES.createDecryptor(key, {
            mode:  cryptoJS.mode.CBC,
            iv: cryptoJS.enc.Hex.parse(iv),
            padding: cryptoJS.pad.Pkcs7,
        });
    }

    private kdf(password: string, salt: cryptoJS.lib.WordArray) {
        return cryptoJS.PBKDF2(password, salt, {
            keySize: this.keyLength, 
            iterations: this.numPBKDF2Iterations, 
            hasher: cryptoJS.algo.SHA256
        });
    }

    public decrypt(encryptedStr: string): string {
        const encrypted = cryptoJS.enc.Hex.parse(encryptedStr);
        return this.decryptor.finalize(encrypted).toString(cryptoJS.enc.Latin1);
    }
}