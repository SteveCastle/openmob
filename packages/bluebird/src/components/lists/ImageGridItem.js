import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import {
  space,
  width,
  fontSize,
  color,
  borders,
  position,
  bottom,
  borderRadius
} from 'styled-system';

const StyledItem = styled.li`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${position}
  ${fontSize}
  ${color}
  box-sizing: border-box;
  min-height:200px;
  transition: 2s;
  cursor: pointer;
  :hover {
    opacity: .7
  }
`;

const StyledImage = styled.div`
  width: 100%;
  height: 100%;
  display: flex;
  align-items: flex-start;
  flex-direction: column;
  justify-content: flex-end;
  flex-direction: column;
  flex-wrap: wrap;
  background-image: url("${props => props.image}");
  background-position: center;
  background-origin: content-box;
  background-repeat: no-repeat;
  background-size: cover;
`;

const StyledTitle = styled.h2`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${fontSize}
  ${position}
  ${bottom}
  ${color}
  text-shadow: 1px 1px 1px #000;
`;

const StyledCaption = styled.p`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${fontSize}
  ${position}
  ${bottom}
  ${color}
  text-shadow: 1px 1px 1px #000;
`;

function ImageGridItem({ uri, width, title, caption, onClick }) {
  return (
    <StyledItem
      p={2}
      position="relative"
      width={[1, width / 12]}
      onClick={onClick}
    >
      <StyledImage image={uri} position={'absolute'} p={3}>
        <StyledTitle color="gray.0" mx={2} my={0}>
          {title}
        </StyledTitle>
        <StyledCaption color="gray.0" mx={2} my={1}>
          {caption}
        </StyledCaption>
      </StyledImage>
    </StyledItem>
  );
}

ImageGridItem.propTypes = {
  uri: PropTypes.string,
  width: PropTypes.number,
  title: PropTypes.string,
  caption: PropTypes.string,
  onClick: PropTypes.func
};
ImageGridItem.defaultProps = { width: 3 };

export default ImageGridItem;
