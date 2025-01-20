 import './searchbar.css'

const Searchbar: React.FC = () => {
    return (
    <div className="search-container">
        <input type="text" className="search-bar" placeholder="ค้นหา..."></input>
    </div>
    );
};

export default Searchbar;