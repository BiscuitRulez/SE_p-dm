import React from 'react';
import './index.css';

type TicketProps = {
  title: string;
  description: string;
  children?: React.ReactNode;
  isCollected: boolean;
  onCollect: () => Promise<void>;
  imageUrl: string;
};

const Ticket: React.FC<TicketProps> = ({
  title,
  description,
  children,
  isCollected,
  onCollect,
  imageUrl
}) => {
  return (
    <div className="ticket">
      <div className="ticket-content">
        {/* Image */}
        <div className="ticket-image">
          <img src={imageUrl} alt={title} className="ticket-image-img" />
        </div>
        
        {/* Content */}
        <div className="ticket-text">
          <div className="ticket-header">
          
            <h2>{title}</h2>
          </div>
          <p>{description}</p>
          {children}
        </div>
        
        {/* Collect Button and Terms */}
        <div className="ticket-action">
          <button
            className="ticket-button"
            onClick={onCollect}
            disabled={isCollected}
          >
            {isCollected ? "เก็บโค้ดแล้ว" : "เก็บโค้ด"}
          </button>
        </div>
      </div>
    </div>
  );
};

export default Ticket;