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
        salt = "dfdf45b433ba4a3082ffd106a688b4c1b3e5efa07e73e208c46178e9abb9d816";
        iv = "e787923d9d0b874eecb45a6677a66f8c";
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

        const plaintext = "PING google.com (216.58.213.78): 56 data bytes";
        const encrypted = cryptoJS.enc.Hex.parse("703575fe8fd11e356b02892173a724e511955f9636119244440b93c8b52bf701997e9318dc0a923671d67dbc1c5492");
        let dec = decryptor.finalize(encrypted).toString(cryptoJS.enc.Latin1);
        console.log(dec)
        
        //let suffix = decryptor.finalize();
        //stdout.write(suffix.toString(cryptoJS.enc.Latin1));
        const encrypted2 = cryptoJS.enc.Hex.parse("686a6de4e0374644df773fb18ade30dbd2c5240fbedfa20c8ca30ebd90e1813f4e19d9e58d24dfe36c19c683b8b550077729267f3fbdef97ce462b65874b");
        let dec2 = decryptor.finalize(encrypted2).toString(cryptoJS.enc.Latin1);
        console.log(dec2);
        
        /*
        const encrypted3 = cryptoJS.enc.Hex.parse("f27e361a971a541bc81756bf9ed66303e05531d86e92ef22202029512ea16e7ce8fd2f021c05b9502b41535ad2d757af4a457a5dde80ff08033e8e250ac1");
        let dec3 = decryptor.process(encrypted3);
        console.log(dec3.toString(cryptoJS.enc.Latin1));

        const encrypted4 = cryptoJS.enc.Hex.parse("b6d93b561dd289727d893b2a5c6c6e3c13506e2cb1289edd542b073cb18371b8af2bc05ff118749e347483d361305d317035dbb2788212f5a4847ba823fe");
        let dec4 = decryptor.finalize(encrypted4);
        console.log(dec4.toString(cryptoJS.enc.Latin1));
        */
        let suffix = decryptor.finalize().toString(cryptoJS.enc.Latin1);
        console.log(suffix);
        console.log("###")
    }

    public decrypt(encrypted: string): string {
        //console.log("Decrypted: ", this.decryptor.process(encrypted).ciphertext.toString());
        return encrypted;
    }
}