import testprofilepic from '../../public/testprofilepic.jpg'

const styles = {
    //rounded profile picture with a size of 64x64px max, showing background color if no picture is available
    avatar: {
        borderRadius: '50%',
        overflow: 'hidden',

        maxHeight: '40px',
        maxWidth: '40px',

        backgroundColor: 'red',
    },

    avatarImage: {
        maxHeight: '40px',
        maxWidth: '40px',
    },

    avatarInitials: {
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',

        height: '40px',
        width: '40px',

        backgroundColor: '#581095',
        color: '#20F9B8',

        fontSize: '20px',
        fontWeight: 'bold',
    }
}

//TODO: this component will temporarily be a placeholder for the users avatar (profile picture or initials)
function avatarContent() {
    if (true) { //TODO: this is temporary, flicking this to test the two different states
        return (
            <div style={styles.avatarInitials}>
               AD
            </div>
        )
    } else {
        return (
            <div>
                <img style={styles.avatarImage} src={testprofilepic} alt="user profile picture"/>
            </div>
        )

    }
}

export default function Avatar() {
    return (
        <div style={styles.avatar}>
            {avatarContent()}
        </div>
    )
}