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
            color={'gray.1'}
            p={2}
            fontSize={3}
            bg={'grey.9'}
            ref={ref}
            onClick={() => setActive(!active)}
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
