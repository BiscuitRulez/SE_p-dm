import { Button } from "antd";

import { useNavigate } from "react-router-dom";




function HomeCodeandPromotion() {

    const navigate = useNavigate();

    const OnCodePage = () => {
        navigate("/code");
    } 


    return (
        <div className="Box">
            <header className="Code">
                <Button
                type="primary"
                onClick={OnCodePage}
                >go to custom code
                </Button>
            </header>
        
        </div>
    );
}

export default HomeCodeandPromotion