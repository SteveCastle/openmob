import React, { useState } from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import { Manager, Reference, Popper } from 'react-popper';

import { space, fontSize, color, borders, borderRadius } from 'styled-system';

const StyledControl = styled.div`
  ${borders}
  ${borderRadius}
  ${space}
  ${fontSize}
  ${color}
  background-color: black;
  cursor: pointer;
`;

const StyledPopper = styled.div`
  ${borders}
  ${borderRadius}
  ${space}
  ${fontSize}
  ${color}
  cursor: pointer;
`;

function Control({ onClick, children }) {
  const [active, setActive] = useState(false);
  return (
    <Manager>
      <Reference>
        {({ ref }) => (
          <StyledControl
            border={0}
            borderRadius={2}
            color={'gray.1'}
            p={2}
            m={1}
            fontSize={2}
            ref={ref}
            onClick={e => {
              e.stopPropagation();
              setActive(!active);
            }}
          >
            {children}
          </StyledControl>
        )}
      </Reference>

      {active && (
        <Popper placement="bottom">
          {({ ref, style, placement, arrowProps }) => (
            <StyledPopper
              ref={ref}
              style={style}
              data-placement={placement}
              bg="yellow.5"
              borderRadius={2}
              p={2}
            >
              <button onClick={onClick}>Confirm</button>
              <div ref={arrowProps.ref} style={arrowProps.style} />
            </StyledPopper>
          )}
        </Popper>
      )}
    </Manager>
  );
}

Control.propTypes = {
  onClick: PropTypes.func
};
Control.defaultProps = {};

export default Control;
