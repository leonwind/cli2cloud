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
                Install the terminal client directly from the source by running
                <div className={styles.codeBlock}>
                    <code className={styles.codeStyle}>
                        $ go install github.com/leonwind/cli2cloud/cli/cli2cloud@latest 
                    </code>
                </div>
            </h6>

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
                    Share and monitor it live from https://cli2cloud.com/4sYe3G <br/>
                    <br/>
                    PING google.com (172.217.22.142): 56 data bytes <br/>
                    64 bytes from 172.217.22.142: icmp_seq=0 ttl=112 time=12.306 ms <br/>
                    64 bytes from 172.217.22.142: icmp_seq=1 ttl=112 time=14.317 ms <br/>
                    ...
                    </code>
                </div>
                <br/>
                and open <code className={styles.codeStyle}>https://cli2cloud.com/{"{your ID}"}</code> on any browser you have.
                It will pipe both your <code className={styles.codeStyle}>Stdout</code> and your <code className={styles.codeStyle}>Stderr</code> output to the web.
            </h6>
            
            <h6 className={styles.docsSubTitle} id={"e2ee"} onClick={() => {highlightElement("e2ee")}}>
                ## End-to-End encryption
            </h6>

            <h6 className={styles.instruction}>
                Use the <code className={styles.codeStyle}>-encrypt {"{password}"}</code> option
                to encrypt your data End-to-End using the <a href="https://en.wikipedia.org/wiki/Block_cipher_mode_of_operation#Cipher_block_chaining_(CBC)">AES CBC Mode</a> { }
                and a 256 bit key generated based on your password using the <a href="https://en.wikipedia.org/wiki/PBKDF2">PBKDF2</a> function.

                <div className={styles.codeBlock}>
                    <code className={styles.codeStyle}>
                    $ ping google.com | cli2cloud -encrypt 1234 <br/>
                    Your client ID: CGYWdD <br/>
                    Share and monitor it live from https://cli2cloud.com/CGYWdD?key=1234<br/>
                    <br/>
                    PING google.com (172.217.22.142): 56 data bytes <br/>
                    64 bytes from 172.217.22.142: icmp_seq=0 ttl=112 time=14.154 ms <br/>
                    64 bytes from 172.217.22.142: icmp_seq=1 ttl=112 time=12.565 ms <br/>
                    ...
                    </code>
                </div>
                <br/>

                To decrypt the data on the web, you need to enter the same password again. 
                The server does not store your password or the hash of it and thus can't validate if your password is either correct or incorrect. 
                You will see complete garbage if you enter a wrong password :)
                <br/>
                <br/>
                Use the option <code className={styles.codeStyle}>-encrypt-random</code> to generate a random secure password with 16 symbols.
                <div className={styles.codeBlock}>
                    <code className={styles.codeStyle}>
                    $ ping google.com | cli2cloud -encrypt-random <br/>
                    Your password: mruI3ubFXTww1QYf <br/>
                    Your client ID: 56xY35 <br/>
                    Share and monitor it live from https://cli2cloud.com/56xY35?key=mruI3ubFXTww1QYf<br/>
                    <br/>
                    PING google.com (142.250.201.174): 56 data bytes <br/>
                    64 bytes from 142.250.201.174: icmp_seq=0 ttl=116 time=3.322 ms <br/>
                    64 bytes from 142.250.201.174: icmp_seq=1 ttl=116 time=2.648 ms <br/>
                    ... 
                    </code>
                </div>
            </h6>
            
            <h3 className={styles.docsTitle} id={"feedback"} onClick={() => {highlightElement("feedback")}}>
                # Feedback
            </h3>
            <h6 className={styles.instruction}>
                The code is open-source available on <a href="https://github.com/leonwind/cli2cloud">GitHub</a>.
                Feel free to open a <a href="https://github.com/leonwind/cli2cloud/issues/new">new Issue</a> regarding any feedback, bugs or feature requests.
            </h6>

            <div className={styles.spaceToBottom}/>
        </div>
    )
}