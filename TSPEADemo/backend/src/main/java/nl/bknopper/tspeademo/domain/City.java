package nl.bknopper.tspeademo.domain;

import lombok.Data;

@Data
public class City {

    public static final double EARTH_RADIUS = 3958.75;

    private String name;
    private double latitude;
    private double longitude;

    public City(String name, double latitude, double longitude) {
        this.name = name;
        this.latitude = latitude;
        this.longitude = longitude;
    }

    public long calculateDistance(City otherCity) {
        double dLat = Math.toRadians(otherCity.latitude - this.latitude);
        double dLng = Math.toRadians(otherCity.longitude - this.longitude);
        double a = Math.sin(dLat / 2) * Math.sin(dLat / 2)
            + Math.cos(Math.toRadians(this.latitude))
            * Math.cos(Math.toRadians(otherCity.latitude))
            * Math.sin(dLng / 2) * Math.sin(dLng / 2);
        double c = 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1 - a));
        long dist = Math.round(EARTH_RADIUS * c * 1.609344);

        return dist;
    }
}
