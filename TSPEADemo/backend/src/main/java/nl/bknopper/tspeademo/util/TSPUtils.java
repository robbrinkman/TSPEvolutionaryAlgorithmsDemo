package nl.bknopper.tspeademo.util;

import nl.bknopper.tspeademo.domain.City;

import java.util.Arrays;
import java.util.Collections;
import java.util.List;

public final class TSPUtils {

    private static final City baseCity = new City("Amsterdam", 52.370216,
	    4.89516);

    private static final City[] cities = new City[] {
	    new City("Athens", 37.975334, 23.736151),
	    new City("Bern", 46.947922, 7.444608),
	    new City("Berlin", 52.519173, 13.406091),
	    new City("Bratislava", 48.146240, 17.107262),
	    new City("Brussels", 50.850342, 4.351710),
	    new City("Bucharest", 44.437710, 26.097366),
	    new City("Budapest", 47.498405, 19.040758),
	    new City("Copenhagen", 55.676098, 12.568337),
	    new City("Dublin", 53.344105, -6.267494),
	    new City("Helsinki", 60.166588, 24.943556),
	    new City("Lisbon", 38.706932, -9.135632),
	    new City("London", 51.508129, -0.128005),
	    new City("Luxemburg", 49.611622, 6.131935),
	    new City("Madrid", 40.416691, -3.700345),
	    new City("Oslo", 59.913868, 10.752245),
	    new City("Prague", 50.087811, 14.420460),
	    new City("Rome", 41.890518, 12.494249),
	    new City("Sofia", 42.696491, 23.326012),
	    new City("Stockholm", 59.328930, 18.064911),
	    new City("Vienna", 48.208176, 16.373819),
	    new City("Warsaw", 52.229675, 21.012230), };

    public static final City[] getCities() {
	return cities;
    }

    public static final List<City> getRandomizedCities() {
	List<City> citiesList = Arrays.asList(Arrays.copyOf(cities,
		cities.length));
	Collections.shuffle(citiesList);
	return citiesList;
    }

    public static City getBaseCity() {
	return baseCity;
    }
}
