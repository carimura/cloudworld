package com.example.fn;

import com.fnproject.fn.api.flow.Flow;
import com.fnproject.fn.api.flow.FlowFuture;
import com.fnproject.fn.api.flow.Flows;
import com.fnproject.fn.api.flow.HttpResponse;

import java.io.Serializable;
import java.util.ArrayList;
import java.util.List;

class ReservationsRequest {
    public String search;
}

class Reservation implements Serializable {
    public String id;
    public String email;
    public String destination;
    public String departing;
    public String returning;

}

class Reservations implements Serializable {
    public List<Reservation> reservations;

}

class BookHotelResult implements Serializable {
    String stuff;
}

class BookFlightResult implements Serializable {
    String stuff;
}

public class FlowFunction {
    public void handleRequest() throws Exception {
        Flow fl = Flows.currentFlow();

        FlowFuture<Reservations> reservations = fl.invokeFunction("./get_reservations", new ReservationsRequest(), Reservations.class);

        FlowFuture<Void> allTasks = reservations.thenCompose((result) -> {
            List<FlowFuture<Void>> results = new ArrayList<>();

            for (Reservation reservation : result.reservations) {
                FlowFuture<HttpResponse> processReservationResult = fl.invokeFunction("./process_reservation", reservation);

                FlowFuture<Void> runBookingsResult = processReservationResult.thenCompose((ignored) -> {
                    FlowFuture<?> bookFlight = fl.invokeFunction("./book_flight", reservation);
                    FlowFuture<?> bookHotel = fl.invokeFunction("./book_hotel", reservation);
                    return fl.allOf(bookFlight,bookHotel);
                });

                FlowFuture<Void> cardDone = runBookingsResult.thenAccept((ignored)->{
                    System.err.println("finished processing " + reservation.email );
                });
                results.add(cardDone);
            }

            return fl.allOf(results.toArray(new FlowFuture[results.size()]));
        });

        allTasks.thenAccept((ignored) -> {
            System.err.println("Flow done");
        });

    }
}