/*
 * NCIM Groep
 * 
 * Created on : 
 * Author     : Niels van Eijck
 * 
 * This class is used for demo's for Evolutionary Algorithms
 *
 */

package nl.bknopper.tspeademo.gui;

import gov.nasa.worldwind.BasicModel;
import gov.nasa.worldwind.awt.WorldWindowGLCanvas;
import gov.nasa.worldwind.geom.Position;
import gov.nasa.worldwind.layers.RenderableLayer;

import java.awt.BorderLayout;
import java.awt.Dimension;
import java.awt.EventQueue;
import java.awt.Image;
import java.lang.reflect.InvocationTargetException;
import java.util.List;

import javax.swing.ImageIcon;
import javax.swing.JFrame;
import javax.swing.JOptionPane;
import javax.swing.JSplitPane;

import nl.bknopper.tspeademo.domain.City;
import nl.bknopper.tspeademo.ea.CandidateSolution;
import nl.bknopper.tspeademo.util.TSPUtils;

public final class TSPEADemo {

	private static final WorldWindowGLCanvas wwd = new WorldWindowGLCanvas();

	private static final AppFrame frame = new AppFrame();

	private AlgorithmUI algorithmUI;

	private double currentBestCandidateSolutionFound;
	private int currentGeneration;

	public TSPEADemo() {
		algorithmUI = new AlgorithmUI(this);
		startView(algorithmUI);
	}

	private static class AppFrame extends JFrame {

		private static final long serialVersionUID = 6954901339876266277L;
		private RenderableLayer layer;

		public AppFrame() {
			initializeWWD();
		}

		private void initializeWWD() {
			wwd.setPreferredSize(new Dimension(900, 960));
			wwd.setModel(new BasicModel());
			wwd.getView().setEyePosition(
					Position.fromDegrees(52.18958, 5.29524, 6e6));
			layer = TSPUtils.createCitiesLayer("Cities");
			wwd.getModel().getLayers().add(layer);
		}

		public void showRoute(List<City> route, double fitness) {

			wwd.getModel().getLayers().remove(layer);
			layer = TSPUtils.createCitiesLayer("Cities");
			wwd.getModel().getLayers().add(layer);
			TSPUtils.buildCitiesLayer(layer, route);
			frame.setTitle("JavaOne TSP Demo. Fitness: " + fitness);
		}

		public void reset() {
			wwd.getModel().getLayers().remove(layer);
			layer = TSPUtils.createCitiesLayer("Cities");
			wwd.getModel().getLayers().add(layer);

			frame.setTitle("JavaOne TSP Demo.");
		}
	}

	public void startView(final AlgorithmUI algorithmUI) {
		try {
			EventQueue.invokeAndWait(new Runnable() {
				public void run() {
					frame.setTitle("JavaOne TSP Demo");
					frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
					frame.setVisible(true);

					JSplitPane splitPane = new JSplitPane(
							JSplitPane.HORIZONTAL_SPLIT);
					splitPane.add(algorithmUI, JSplitPane.LEFT);
					splitPane.add(wwd, JSplitPane.RIGHT);
					frame.add(splitPane, BorderLayout.CENTER);
					frame.pack();
				}
			});
		} catch (InterruptedException ex) {
			// ignore
		} catch (InvocationTargetException ex) {
			// ignore
		}
	}

	public void showLastGeneration(final CandidateSolution candidateSolution,
			final int generation) {
		EventQueue.invokeLater(new Runnable() {
			public void run() {
				frame.showRoute(candidateSolution.getRoute(),
						candidateSolution.getFitness());
				algorithmUI.showAlgorithmInfo(candidateSolution.getFitness(),
						generation);

			}
		});

		this.currentBestCandidateSolutionFound = candidateSolution.getFitness();
		this.currentGeneration = generation;
	}

	public void enableStartButton(boolean b) {
		algorithmUI.enableStartButton(b);
	}

	public void enableStopButton(boolean b) {
		algorithmUI.enableStopButton(b);
	}

	public void reset() {
		enableStartButton(true);
		enableStopButton(false);

		EventQueue.invokeLater(new Runnable() {
			public void run() {
				frame.reset();
				algorithmUI.showAlgorithmInfo(Double.NaN, 0);
			}
		});
	}

	public void done() {
		enableStartButton(true);
		enableStopButton(false);
		ImageIcon icon = new ImageIcon(getClass().getClassLoader().getResource(
				"images/done.png"));

		// resize image in icon
		Image img = icon.getImage();
		Image newimg = img.getScaledInstance(150, 150,
				java.awt.Image.SCALE_SMOOTH);
		icon = new ImageIcon(newimg);

		JOptionPane
				.showMessageDialog(
						frame,
						"<html>"
								+ "<p style='margin:20px; margin-left:0px; margin-right:40px;'>"
								+ "<span style='font-size:1.5em'>The Evolutionary Algorithm has finished.<br /><br />"
								+ "Best Candidate Solution Found: "
								+ currentBestCandidateSolutionFound + "<br />"
								+ "In number of generations: "
								+ currentGeneration
								+ "<br /><br /></span></p></html>", "Done!",
						JOptionPane.OK_OPTION, icon);

	}

	public static void main(String[] args) {
		new TSPEADemo();
	}
}