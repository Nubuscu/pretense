/**
 * Styles for what goes inside cytoscape.
 * node and edge weights and styles, etc.
 */
export default [
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
      'background-color': '#458588', // blue-dim
      'border-color': '#83a598', // blue
      'border-width': '1',
      'color': '#d5c4a1', // fg2
    }
  },
  {
    selector: 'node[isTopic]',
    style: {
      'width': '40',
      'height': '40',
      'shape': 'hexagon',
      'background-color': '#b16286', // purple-dim
      'border-color': '#d3869b', // purple
    }
  },
  {
    selector: 'edge',
    style: {
      'curve-style': 'bezier',
      'color': '#7c6f64',
      'text-background-color': '#ffffff',
      'text-background-opacity': '1',
      'text-background-padding': '3',
      'width': '3',
      'target-arrow-shape': 'none',
      'line-color': '#7c6f64',  // bg4
      'target-arrow-color': '#7c6f64',
      'font-weight': 'bold'
    }
  },
  {
    selector: 'edge[label]',
    style: {
      'content': `data(label)`,
    }
  },
  {
    selector: 'edge.label',
    style: {
      'line-color': 'orange',
      'target-arrow-color': 'orange'
    }
  }
]