import * as cryptoJS from "crypto-js"

export class DecryptionService {
    keyLength: number = 8;
    numPBKDF2Iterations: number = 1024;

    key: cryptoJS.lib.WordArray; // crypto.WordArray
    decryptor: any; // crypto.Cipher

    private kdf(password: string, salt: cryptoJS.lib.WordArray) {
        return cryptoJS.PBKDF2(password, salt, {
            keySize: this.keyLength, 
            iterations: this.numPBKDF2Iterations, 
            hasher: cryptoJS.algo.SHA256
        });
    }

    public constructor(password: string, salt: string, iv: string) {
        salt = "836440f899b9050ffeed785004bd095d63bc6c6d22264f6da562d6d142199e36";
        iv = "bf8ecd752ca283b2b40d8a0bf30b6176";
        this.key = this.kdf(password, cryptoJS.enc.Hex.parse(salt))
        console.log("Keys:");
        console.log(this.key.toString());
        //console.log("0027f656ef128fd9b5f563336f50fad44d78435f6739c3b4b656dbdc20ca8454");
        console.log("####");

        const decryptor = cryptoJS.algo.AES.createDecryptor(this.key, {
            mode:  cryptoJS.mode.CBC,
            iv: cryptoJS.enc.Hex.parse(iv),
            padding: cryptoJS.pad.Pkcs7,
        });
        
        const encrypted = cryptoJS.enc.Hex.parse("bb1f27dae3afa314e0c8a3e8ce06183399f208618ada31c7979e8a047e2cce11f29ab9e0d4df6aef25efefd1124a0ad5");
        let dec = decryptor.finalize(encrypted).toString(cryptoJS.enc.Latin1);
        console.log(dec)
        
        const encrypted2 = cryptoJS.enc.Hex.parse("540810c10e1842500282c84749f71e287cdf9bf19699844f0329523771362fc4b6a1ad91644acfee28b252937ecd915cf4684743aa133f25b1a6c4f71a1ce74e");
        let dec2 = decryptor.finalize(encrypted2).toString(cryptoJS.enc.Latin1);
        console.log(dec2);
        
        const encrypted3 = cryptoJS.enc.Hex.parse("fe5ff7383fb3190055aa9138c78d592775b1c39652122b1a16cca169bef168948869784bd3c9b9fb4f1799efa80bc6a19f782f62cf4f627ae184a45bde298731");
        let dec3 = decryptor.finalize(encrypted3);
        console.log(dec3.toString(cryptoJS.enc.Latin1));

        const encrypted4 = cryptoJS.enc.Hex.parse("e8d5e110a662e141a0493cf5d1dd21762e8b69269beeb1ef1cfb6597614bb33b8284d11693fb90a3227cabe0fc232e76a6fd2f7f26692c14f86ac169769fdf34");
        let dec4 = decryptor.finalize(encrypted4);
        console.log(dec4.toString(cryptoJS.enc.Latin1));
    }

    public decrypt(encrypted: string): string {
        //console.log("Decrypted: ", this.decryptor.process(encrypted).ciphertext.toString());
        return encrypted;
    }
}