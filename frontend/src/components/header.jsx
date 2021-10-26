import Logo from "./logo.jsx";
import "../css/header.css";
import SearchBar from "./searchBar.jsx";
import ProfileImage from "./profileImage.jsx";
import dummyProfile from "../media/mochi.jpeg";

function Header({ showSearch, showProfile}) {
    return (
        <div className="d-flex flex-row justify-content-between">
            <div className="logo-spacing">
                <Logo />
            </div>
            {showSearch ? <SearchBar showFilterBtn={true}/> : <></>}
            {showProfile ? <ProfileImage url={dummyProfile} /> : <></>}
        </div>
    );
}

export default Header;
