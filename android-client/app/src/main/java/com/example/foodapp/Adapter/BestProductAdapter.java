package com.example.foodapp.Adapter;

import android.annotation.SuppressLint;
import android.content.Context;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ImageView;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.recyclerview.widget.RecyclerView;

import com.bumptech.glide.Glide;
import com.bumptech.glide.load.resource.bitmap.CenterCrop;
import com.bumptech.glide.load.resource.bitmap.RoundedCorners;
import com.example.foodapp.Domain.Product.ProductResponse;
import com.example.foodapp.R;

import java.util.ArrayList;

public class BestProductAdapter extends RecyclerView.Adapter<BestProductAdapter.viewholder> {
    ArrayList<ProductResponse> items;
    Context context;

    public BestProductAdapter(ArrayList<ProductResponse> items) {
        this.items = items;
    }

    @NonNull
    @Override
    public BestProductAdapter.viewholder onCreateViewHolder(@NonNull ViewGroup parent, int viewType) {
        context = parent.getContext();
        View inflate = LayoutInflater.from(parent.getContext()).inflate(R.layout.viewholder_bestproduct, parent, false);
        return new viewholder(inflate);
    }

    @Override
    public void onBindViewHolder(@NonNull BestProductAdapter.viewholder holder, @SuppressLint("RecyclerView") int position) {
        holder.nameText.setText(items.get(position).getName());
        holder.priceText.setText(items.get(position).getPrice()+"BYN");
        holder.timeText.setText("20min");

        Glide.with(context)
                .load(items.get(position).getImageURL())
                .transform(new CenterCrop(), new RoundedCorners(30))
                .into(holder.pic);
    }

    @Override
    public int getItemCount() {
        return items.size();
    }

    public static class viewholder extends RecyclerView.ViewHolder {
        TextView nameText, priceText, timeText;
        ImageView pic;
        public viewholder(@NonNull View itemView) {
            super(itemView);
            nameText = itemView.findViewById(R.id.nameText);
            timeText = itemView.findViewById(R.id.timeText);
            priceText = itemView.findViewById(R.id.priceText);
            pic = itemView.findViewById(R.id.pic);
        }
    }
}