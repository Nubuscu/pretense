/**
 * Styles for what goes inside cytoscape.
 * node and edge weights and styles, etc.
 */

export const colours = {
  // https://materialtheme.arcsine.dev/
  primary: '#512da8',
  primaryDark: '#371c8d',
  primaryLight: '#cbc0e5',
  edgeDefault: '#797979',
  edgeSelected: '#cbc0e5',
  text: '#ffffff',
  bg: '#2c2c2c',
}
export const graphStyles = [
  {
    selector: 'node',
    style: {
      'width': '20',
      'height': '20',
      'font-size': '10',
      'font-weight': 'bold',
      'content': `data(label)`,
      'text-valign': 'center',
      'text-wrap': 'wrap',
      'text-max-width': '140',
      'background-color': colours.primary,
      'border-color': colours.primary,
      'border-width': '1',
      'color': colours.text,
      'layoutDimensions': {
        'nodeDimensionsIncludeLabels': true,

      }
    }
  },
  {
    selector: 'node[isTopic]',
    style: {
      'width': '40',
      'height': '40',
      'shape': 'hexagon',
      'background-color': colours.primaryDark,
      'border-color': colours.primaryDark,
    }
  },
  {
    selector: 'edge',
    style: {
      'curve-style': 'bezier',
      'color': colours.primaryLight,
      'text-background-color': '#ffffff',
      'text-background-opacity': '1',
      'text-background-padding': '3',
      'width': '3',
      'target-arrow-shape': 'none',
      'line-color': colours.edgeDefault,
      'font-weight': 'bold'
    }
  },
  {
    selector: 'edge[source ^= "topic"][target ^= "album"]',
    style: {
      'line-color': colours.primaryLight,
    }
  },
]