import styles from "../styles/Documentation.module.css";

export const Documentation = () => {

    const highlightElement = (id: string) => {
        window.location.hash = id;
    }

    return (
        <div className={styles.body}>
            <h3 className={styles.docsTitle} id={"installation"} onClick={() => {highlightElement("installation")}}>
                # Installation     
            </h3>

            <h6 className={styles.instruction}>
                Install the client by running { }
                <div className={styles.codeBlock}>
                    <code className={styles.codeStyle}>
                    $ go get github.com/leonwind/cli2cloud/cli/cli2cloud 
                    </code>
                </div>
            </h6>

            <br/>

            <h3 className={styles.docsTitle} id={"examples"} onClick={() => {highlightElement("examples")}}>
                # Examples 
            </h3>
            
            <h6 className={styles.docsSubTitle} id={"normal-usage"} onClick={() => {highlightElement("normal-usage")}}>
                ## Normal usage
            </h6>

            <h6 className={styles.instruction}>
                To just pipe your terminal output from any arbitrary command, run
                <div className={styles.codeBlock}>
                    <code className={styles.codeStyle}>
                    $ ping google.com | cli2cloud <br/>
                    Your client ID: 4sYe3G <br/>
                    Share and monitor it live from cli2cloud.com/4sYe3G <br/>
                    <br/>
                    PING google.com (172.217.22.142): 56 data bytes <br/>
                    64 bytes from 172.217.22.142: icmp_seq=0 ttl=112 time=12.306 ms <br/>
                    64 bytes from 172.217.22.142: icmp_seq=1 ttl=112 time=14.317 ms <br/>
                    ...
                    </code>
                </div>
            </h6>
            
            <h6 className={styles.docsSubTitle} id={"e2ee"} onClick={() => {highlightElement("e2ee")}}>
                ## End-to-End encryption
            </h6>
            <h6 className={styles.instruction}>
                Use the option <code className={styles.codeStyle}>-encrypt=password</code> {  }
                to encrypt your data End-to-End using the <a href="https://en.wikipedia.org/wiki/Block_cipher_mode_of_operation#Cipher_block_chaining_(CBC)">AES CBC mode</a> { }
                and a 256 Bit key generated based on your password:

                <div className={styles.codeBlock}>
                    <code className={styles.codeStyle}>
                    $ ping google.com | cli2cloud -encrypt=1234<br/>
                    Your client ID: CGYWdD <br/>
                    Share and monitor it live from cli2cloud.com/CGYWdD <br/>
                    <br/>
                    PING google.com (172.217.22.142): 56 data bytes <br/>
                    64 bytes from 172.217.22.142: icmp_seq=0 ttl=112 time=14.154 ms <br/>
                    64 bytes from 172.217.22.142: icmp_seq=1 ttl=112 time=12.565 ms <br/>
                    ...
                    </code>
                </div>
            </h6>
            
            <h3 className={styles.docsTitle} id={"feedback"} onClick={() => {highlightElement("feedback")}}>
                # Feedback
            </h3>
            <h6 className={styles.instruction}>
                The code is open-source available on <a href="https://github.com/leonwind/cli2cloud">GitHub</a>. <br/>
                Feel free to open a <a href="https://github.com/leonwind/cli2cloud/issues/new">new Issue</a> regarding any feedback or bugs.
            </h6>

            <div className={styles.spaceToBottom}/>
        </div>
    )
}